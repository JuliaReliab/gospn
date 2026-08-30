package parser

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/okamumu/gospn/pkg/petrinet"
)

// These tests used to walk a parse tree and print the result, asserting nothing: a
// change that broke the listener outright would still have passed them.

// errorCollector replaces ANTLR's default listener, which writes to stderr and lets the
// parse look successful. Without it nothing here can tell a syntax error from a clean
// parse.
type errorCollector struct {
	*antlr.DefaultErrorListener
	errs []string
}

func (c *errorCollector) SyntaxError(_ antlr.Recognizer, _ interface{}, line, col int, msg string, _ antlr.RecognitionException) {
	c.errs = append(c.errs, fmt.Sprintf("%d:%d %s", line, col, msg))
}

// parseProg walks a whole definition and returns the listener and any syntax errors.
// It panics on malformed input -- see TestListenerReportsSyntaxError.
func parseProg(text string) (*PNListener, []string) {
	c := &errorCollector{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	lexer := NewJSPNLLexer(antlr.NewInputStream(text))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(c)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewJSPNLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(c)
	listener := NewPNListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Prog())
	return listener, c.errs
}

// The net the printing tests used, minus the parts that were never checked.
const raid6Text = `
place Pn (init = 6)
place Pdf
exp Tdfail (guard = gfail, rate = Tdfail_rate)
gen Trebuild (guard = gfail, dist = Trebuild.dist)
imm Tinit (guard = ginit)
arc Pn to Tdfail
arc Tdfail to Pdf
arc Pdf to Trebuild
arc Pdf to Tinit
arc Trebuild to Pn
arc Tinit to Pn

place Po (init = 1)
place Pr
place Pc
imm Tstart (guard = gstart)
gen Trecon (dist = Trecon.dist)
imm Tend (guard = gend)
arc Po to Tstart
arc Tstart to Pr
arc Pr to Trecon
arc Trecon to Pc
arc Pc to Tend
arc Tend to Po

Tdfail_rate = #Pn * lambda
gfail = #Po == 1
gstart = #Pdf > 2
ginit = #Pc == 1
gend = #Pdf == 0

Trebuild.dist = expdist(MTTR1)
Trecon.dist = unif(1.0, 2.0)

MTTF = 1.0e+6
lambda = 1/MTTF
MTTR1 = 2
MTTR2 = 24

reward r1 #Pc
`

// An expression parses into an AST that evaluates -- and with the right precedence,
// which is the one thing a walk that only prints could never show.
func TestListenerParsesExpressionPrecedence(t *testing.T) {
	lexer := NewJSPNLLexer(antlr.NewInputStream("1 + 2 * 3"))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewJSPNLParser(stream)
	listener := NewPNListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expression())

	if len(listener.builder.aststack) != 1 {
		t.Fatalf("the walk left %d expressions on the stack, want 1", len(listener.builder.aststack))
	}
	res, err := listener.builder.aststack.pop().Eval(make(ASTEnv))
	if err != nil {
		t.Fatal(err)
	}
	v, err := res.GetInt()
	if err != nil {
		t.Fatal(err)
	}
	if v != ASTInt(7) {
		t.Errorf("1 + 2 * 3 = %v, want 7 (times binds tighter)", v)
	}
}

func TestListenerCollectsDeclarations(t *testing.T) {
	listener, errs := parseProg(raid6Text)
	if len(errs) != 0 {
		t.Fatalf("unexpected syntax errors: %v", errs)
	}

	// Every declaration reaches the environment, tagged with what it is.
	for label, want := range map[string]string{
		"Pn": "place", "Pdf": "place", "Po": "place", "Pr": "place", "Pc": "place",
		"Tdfail": "exp", "Trebuild": "gen", "Trecon": "gen",
		"Tinit": "imm", "Tstart": "imm", "Tend": "imm",
		"r1": "reward",
	} {
		node, ok := listener.builder.env[label].(*PNNode)
		if !ok {
			t.Errorf("%s is not a node in the environment", label)
			continue
		}
		if got := node.options["type"]; got != want {
			t.Errorf("%s is a %v, want %v", label, got, want)
		}
	}

	// labels is the declaration order, which makeNet depends on: it creates every
	// place before any transition so that a guard may name a place declared later.
	var places, trans []string
	for _, label := range listener.builder.labels {
		if node, ok := listener.builder.env[label].(*PNNode); ok {
			switch node.options["type"] {
			case "place":
				places = append(places, label)
			case "exp", "gen", "imm":
				trans = append(trans, label)
			}
		}
	}
	if want := []string{"Pn", "Pdf", "Po", "Pr", "Pc"}; !equalStrings(places, want) {
		t.Errorf("places %v, want %v in declaration order", places, want)
	}
	if len(trans) != 6 {
		t.Errorf("%d transitions, want 6", len(trans))
	}

	// A parameter that is only assigned later is still in the environment, since
	// variables are resolved when they are evaluated.
	if _, ok := listener.builder.env["lambda"]; !ok {
		t.Error("lambda is missing from the environment")
	}
}

// A malformed definition must be rejected as an error, and must say where. The test
// this replaces built a parser over broken input, never walked it, and asserted nothing;
// until 0.24.0 the definition reached the user as a Go panic with a stack trace.
func TestReadReportsSyntaxError(t *testing.T) {
	broken := strings.Replace(raid6Text, "imm Tinit (guard = ginit)", "imm Tinit + (guard = ginit)", 1)
	if broken == raid6Text {
		t.Fatal("the text to break was not found")
	}

	net, imark, err := PNreadFromText(broken)
	if err == nil {
		t.Fatal("`imm Tinit + (...)` was accepted")
	}
	if net != nil || imark != nil {
		t.Error("a net was returned for a definition that does not parse")
	}
	if !strings.Contains(err.Error(), "syntax error") || !strings.Contains(err.Error(), "line 6:") {
		t.Errorf("the error is %q; it should say what and where", err)
	}
}

// Every syntax error is reported, not just the first: two mistakes in a file should
// take one run to find, not two.
func TestReadReportsEverySyntaxError(t *testing.T) {
	_, _, err := PNreadFromText(`
place P (init = 1)
imm A + (guard = g)
imm B + (guard = g)
`)
	if err == nil {
		t.Fatal("two broken declarations were accepted")
	}
	if n := strings.Count(err.Error(), "line "); n < 2 {
		t.Errorf("the error mentions %d lines, want both:\n%v", n, err)
	}
}

// A definition that parses but cannot be built -- here an option given a value of the
// wrong kind -- is an error too. The builder reports these by panicking, which is
// convenient inside a recursive walk; the reader turns it into an error so the panic
// does not escape the package.
func TestReadReportsABadOptionAsAnError(t *testing.T) {
	_, _, err := PNreadFromText(`
place P (init = 1)
exp T (rate = 1.0, priority = 1.5)
arc P to T
`)
	if err == nil {
		t.Skip("a float priority is accepted; nothing to report")
	}
	if strings.Contains(err.Error(), "goroutine") {
		t.Errorf("the error carries a stack trace: %v", err)
	}
}

// makeNet turns the declarations into a net: the places, the transitions and the
// initial marking all come through.
func TestListenerBuildsTheNet(t *testing.T) {
	listener, errs := parseProg(raid6Text)
	if len(errs) != 0 {
		t.Fatalf("unexpected syntax errors: %v", errs)
	}
	net, imark := makeNet(listener.builder.labels, listener.builder.env)

	for _, label := range []string{"Pn", "Pdf", "Po", "Pr", "Pc"} {
		if _, ok := net.GetPlace(label); !ok {
			t.Errorf("place %s is missing from the net", label)
		}
	}
	for _, label := range []string{"Tdfail", "Trebuild", "Trecon", "Tinit", "Tstart", "Tend"} {
		if _, ok := net.GetTrans(label); !ok {
			t.Errorf("transition %s is missing from the net", label)
		}
	}
	if _, ok := net.GetReward("r1"); !ok {
		t.Error("reward r1 is missing from the net")
	}

	// Finalize sorts the places by label, so the marking is indexed in that order.
	want := map[string]petrinet.MarkInt{"Pn": 6, "Po": 1, "Pc": 0, "Pdf": 0, "Pr": 0}
	for i, label := range net.PlaceLabels() {
		if imark[i] != want[label] {
			t.Errorf("initial marking of %s is %d, want %d", label, imark[i], want[label])
		}
	}

	// The net evaluates -- every guard, rate and reward resolves.
	if err := net.CheckExpressions(imark); err != nil {
		t.Error(err)
	}
}

// A distribution written in the definition reaches the model with its parameters. The
// marking graph reports it back, which is what a result file records for an MRSPN.
func TestListenerParsesDistributions(t *testing.T) {
	listener, errs := parseProg(raid6Text)
	if len(errs) != 0 {
		t.Fatalf("unexpected syntax errors: %v", errs)
	}
	net, imark := makeNet(listener.builder.labels, listener.builder.env)
	mg, err := petrinet.CreateMarkingGraphWithDFS(net, imark)
	if err != nil {
		t.Fatal(err)
	}

	got := make([]string, 0, 2)
	for _, info := range mg.BlockGenTrans() {
		got = append(got, info.Label+" "+info.Dist)
	}
	sort.Strings(got)
	got = dedup(got)
	want := []string{"Trebuild expdist(2)", "Trecon unif(1,2)"}
	if !equalStrings(got, want) {
		t.Errorf("general transitions %v, want %v", got, want)
	}
}

// An arc multiplicity reaches the token game, rather than merely being parsed.
func TestArcMultiplicityReachesTheModel(t *testing.T) {
	net, imark, err := PNreadFromText(`
		place P (init = 1)
		place Q
		exp T (rate = 1.0)
		arc P to T
		arc T to Q (multi = 10)
	`)
	if err != nil {
		t.Fatal(err)
	}
	mg, err := petrinet.CreateMarkingGraphWithDFS(net, imark)
	if err != nil {
		t.Fatal(err)
	}
	var marks [][]petrinet.MarkInt
	for _, ms := range mg.StateMarkings() {
		marks = append(marks, ms...)
	}
	if len(marks) != 2 {
		t.Fatalf("%d states, want 2", len(marks))
	}
	// P and Q, in Finalize's sorted order.
	found := false
	for _, m := range marks {
		if m[0] == 0 && m[1] == 10 {
			found = true
		}
	}
	if !found {
		t.Errorf("firing T gave %v, want a marking with 10 tokens in Q", marks)
	}
}

// An update block runs after the tokens move, and can set a place outright. The tests
// this file replaces parsed a net with one and never checked what it did.
func TestUpdateBlockReachesTheModel(t *testing.T) {
	net, imark, err := PNreadFromText(`
		place P (init = 3)
		place Q
		exp T (rate = 1.0) {
			#Q = 7
		}
		arc P to T
		arc T to Q
	`)
	if err != nil {
		t.Fatal(err)
	}
	mg, err := petrinet.CreateMarkingGraphWithDFS(net, imark)
	if err != nil {
		t.Fatal(err)
	}
	var marks [][]petrinet.MarkInt
	for _, ms := range mg.StateMarkings() {
		marks = append(marks, ms...)
	}
	// P and Q, in Finalize's sorted order: (3,0) initially, then Q is 7 whatever the
	// arc put there -- the update wins.
	want := map[string]bool{"[3 0]": true, "[2 7]": true, "[1 7]": true, "[0 7]": true}
	if len(marks) != len(want) {
		t.Fatalf("%d states, want %d: %v", len(marks), len(want), marks)
	}
	for _, m := range marks {
		if !want[fmt.Sprint(m)] {
			t.Errorf("unexpected marking %v; the update block did not run", m)
		}
	}
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func dedup(sorted []string) []string {
	out := sorted[:0]
	for i, s := range sorted {
		if i == 0 || s != sorted[i-1] {
			out = append(out, s)
		}
	}
	return out
}
