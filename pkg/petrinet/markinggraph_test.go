package petrinet

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// These tests used to build a marking graph, print it, and assert nothing: the DOT
// writers and the transition matrices were exercised and never checked.

// probeNet is p1 -> t1 -> p3 -> t2 -> {p1, p2}, with t2 as an EXP, IMM or GEN
// transition. t1 consumes one token from p1 and one from p2 and puts one in p3; t2
// takes it back out and returns one token to each. Starting from (10, 1, 1) that gives
// three markings -- (10,1,1), (9,0,2), (11,2,0) -- with t1 disabled in the second (p2 is
// empty) and t2 disabled in the third (p3 is empty).
func probeNet(t *testing.T, t2kind string) (*Net, []MarkInt) {
	t.Helper()
	net := NewNet()
	p1 := net.NewPlace("p1", 100)
	p2 := net.NewPlace("p2", 100)
	p3 := net.NewPlace("p3", 100)
	t1 := net.NewExpTrans("t1", 0, true, 1)
	var t2 transInterface
	switch t2kind {
	case "exp":
		t2 = net.NewExpTrans("t2", 0, true, 1)
	case "imm":
		t2 = net.NewImmTrans("t2", 0, true, 1)
	case "gen":
		t2 = net.NewGenTrans("t2", 0, true, NewDistribution("exponential", 1.0), GenTransPolicyPRD)
	default:
		t.Fatalf("unknown transition kind %q", t2kind)
	}
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()
	return net, []MarkInt{10, 1, 1}
}

func probeGraph(t *testing.T, t2kind string) *MarkingGraph {
	t.Helper()
	net, m0 := probeNet(t, t2kind)
	mg, err := CreateMarkingGraphWithDFS(net, m0)
	if err != nil {
		t.Fatal(err)
	}
	return mg
}

func matrixNames(mg *MarkingGraph) []string {
	exp, imm, gen := mg.TransMatrix()
	grouplabel, translabel := mg.GroupLabels(), mg.TransLabels()
	var names []string
	for _, mats := range []map[GroupTrans]*CSC{exp, imm, gen} {
		for tr := range mats {
			names = append(names, fmt.Sprintf("%s%s%s", grouplabel[tr.src], grouplabel[tr.dest], translabel[tr]))
		}
	}
	sort.Strings(names)
	return names
}

func stateCount(mg *MarkingGraph) int {
	n := 0
	for _, ms := range mg.StateMarkings() {
		n += len(ms)
	}
	return n
}

// The kind of the second transition decides how the markings are grouped, which is what
// the matrix names record.
func TestTransMatrixNamesFollowTheTransitionKind(t *testing.T) {
	for _, tc := range []struct {
		kind   string
		states int
		groups int
		names  []string
		why    string
	}{
		{"exp", 3, 1, []string{"G0G0E"},
			"a pure-EXP net has one group and one matrix, the generator itself"},
		{"imm", 2, 2, []string{"G0G0E", "G0I0E", "I0G0I"},
			"an immediate t2 makes (10,1,1) vanishing: it fires the moment p3 has a token"},
		{"gen", 3, 2, []string{"G0G0E", "G0G1E", "G1G0P0", "G1G1E", "G1G1P0"},
			"a general t2 splits the markings by whether it is aging, and adds P0 jump blocks"},
	} {
		tc := tc
		t.Run(tc.kind, func(t *testing.T) {
			mg := probeGraph(t, tc.kind)
			if got := stateCount(mg); got != tc.states {
				t.Errorf("%d states, want %d (%s)", got, tc.states, tc.why)
			}
			if got := len(mg.groups); got != tc.groups {
				t.Errorf("%d groups, want %d", got, tc.groups)
			}
			if got := matrixNames(mg); !equalStrs(got, tc.names) {
				t.Errorf("matrices %v, want %v", got, tc.names)
			}
		})
	}
}

// The initial marking is in exactly one group, at exactly one position.
func TestInitVectorMarksOneState(t *testing.T) {
	for _, kind := range []string{"exp", "imm", "gen"} {
		kind := kind
		t.Run(kind, func(t *testing.T) {
			mg := probeGraph(t, kind)
			ones, total := 0, 0.0
			for _, v := range mg.InitVector() {
				for _, x := range v {
					total += x
					if x == 1 {
						ones++
					}
				}
			}
			if ones != 1 || total != 1 {
				t.Errorf("the initial vectors hold %d ones summing to %v, want one 1", ones, total)
			}
		})
	}
}

var dotNode = regexp.MustCompile(`^"([^"]+)" \[`)
var dotEdge = regexp.MustCompile(`^"([^"]+)"->"([^"]+)" \[`)

// Node names are the group and the row within it, not addresses. Before this the same
// net produced a different file on every run, so two drawings could not be compared and
// a generated diagram under version control changed every line every time.
func TestMarkDotIsDeterministicAndNamesNodesStably(t *testing.T) {
	for _, kind := range []string{"exp", "imm", "gen"} {
		kind := kind
		t.Run(kind, func(t *testing.T) {
			draw := func(f func(*MarkingGraph, *bytes.Buffer)) string {
				var buf bytes.Buffer
				mg := probeGraph(t, kind)
				f(mg, &buf)
				return buf.String()
			}
			for name, f := range map[string]func(*MarkingGraph, *bytes.Buffer){
				"ToMarkDot":                  func(mg *MarkingGraph, b *bytes.Buffer) { mg.ToMarkDot(b) },
				"ToMarkDotWithLabel":         func(mg *MarkingGraph, b *bytes.Buffer) { mg.ToMarkDotWithLabel(b) },
				"ToMarkDotWithLabelAndGroup": func(mg *MarkingGraph, b *bytes.Buffer) { mg.ToMarkDotWithLabelAndGroup(b) },
				"ToGroupMarkDot":             func(mg *MarkingGraph, b *bytes.Buffer) { mg.ToGroupMarkDot(b) },
			} {
				got := draw(f)
				if again := draw(f); got != again {
					t.Errorf("%s is not reproducible", name)
				}
				if strings.Contains(got, "0x") {
					t.Errorf("%s names a node by its address:\n%s", name, got)
				}
				checkDotIsClosed(t, name, got)
			}
		})
	}
}

// Every edge has to reach nodes the file declares, which is what naming them by index
// rather than by address could get wrong.
func checkDotIsClosed(t *testing.T, name, dot string) {
	t.Helper()
	declared := map[string]bool{}
	var edges [][2]string
	for _, line := range strings.Split(dot, "\n") {
		if m := dotEdge.FindStringSubmatch(line); m != nil {
			edges = append(edges, [2]string{m[1], m[2]})
		} else if m := dotNode.FindStringSubmatch(line); m != nil {
			declared[m[1]] = true
		}
	}
	if len(declared) == 0 || len(edges) == 0 {
		t.Errorf("%s: %d nodes and %d edges", name, len(declared), len(edges))
	}
	for _, e := range edges {
		for _, n := range e {
			if !declared[n] {
				t.Errorf("%s: edge to undeclared node %q", name, n)
			}
		}
	}
}

func equalStrs(a, b []string) bool {
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
