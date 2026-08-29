package parser

import (
	"testing"

	"github.com/okamumu/gospn/pkg/petrinet"
)

// Each guard carries a trailing `&& #p >= 0`, which is always true and does not
// change what is being tested. It is there because createImmTrans/createExpTrans
// render a guard to a string with a markless Eval, and a guard that is constant at
// parse time evaluates to a bool instead and makes that rendering fail -- a
// pre-existing limitation, unrelated to compilation.
//
// The compiled closures have to reproduce the interpreter's arithmetic exactly,
// including the parts of it that are surprising. Each case here is a guard whose
// truth depends on one such rule, checked through the exported IsEnabled; with
// CheckCompiled set, the interpreter runs alongside and any divergence panics, so
// each case pins the rule and the agreement at once.
func TestCompiledArithmetic(t *testing.T) {
	cases := []struct {
		name  string
		guard string
		want  bool
		why   string
	}{
		{"int div yields float", "(1/2) > 0.4", true,
			"ASTInt.div returns float64(a)/float64(b); integer division would give 0"},
		{"int div is not truncated", "(3/2) == 1.5", true, ""},
		{"integer div operator truncates", "(3 div 2) == 1", true,
			"`div` is the integer division, unlike `/`"},
		{"int compared to float promotes", "1 == 1.0", true, ""},
		{"unary minus on int", "(-3) < 0", true, ""},
		{"not", "!(1 > 2)", true, ""},
		{"and", "(1 < 2) && (2 < 3)", true, ""},
		{"or short of the left", "(1 > 2) || (2 < 3)", true,
			"the right operand decides; an evaluator that dropped it would say false"},
		{"ifelse picks the second", "ifelse(1 < 2, 1, 0) == 1", true, ""},
		{"ifelse picks the third", "ifelse(1 > 2, 1, 0) == 0", true, ""},
		{"ifelse over floats", "ifelse(1 > 2, 1.5, 2.5) == 2.5", true, ""},
		{"place token count", "#p == 3", true, ""},
		{"place arithmetic", "(#p + #q) == 3", true, ""},
		{"variable is inlined", "v == 7", true, ""},
		{"variable of a variable", "w == 7", true, ""},
		{"marking-dependent variable", "vp == 3", true, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			CheckCompiled = true
			defer func() { CheckCompiled = false }()
			text := `
place p (init = 3)
place q (init = 0)
exp t (guard = (` + c.guard + `) && #p >= 0, rate = 1)
v = 7
w = v
vp = #p
`
			net, mark := PNreadFromText(text)
			tr, ok := net.GetTrans("t")
			if !ok {
				t.Fatal("transition t not found")
			}
			got := tr.IsEnabled(net, mark) == petrinet.ENABLE
			if got != c.want {
				t.Errorf("guard %q: got %v, want %v %s", c.guard, got, c.want, c.why)
			}
		})
	}
}

// Everything above must actually have been compiled; if it all fell back the test
// would pass while checking nothing about this package.
func TestCompiledArithmeticIsCompiled(t *testing.T) {
	CompileStatsReset()
	PNreadFromText(`
place p (init = 3)
exp t (guard = ((1/2) > 0.4) && #p >= 0, rate = #p * 2.0)
`)
	if CompileStats.GuardCompiled != 1 || CompileStats.RateCompiled != 1 {
		t.Errorf("expected the guard and the rate to compile, got %+v", CompileStats)
	}
}

// An expression this package cannot compile must fall back to the interpreter
// rather than fail to build. Evaluating this particular one still panics, because
// an undefined variable evaluates to its own name as a string and no guard can be
// made of that -- but that is the interpreter's long-standing behaviour and is what
// falling back is supposed to preserve.
func TestUncompilableFallsBack(t *testing.T) {
	CompileStatsReset()
	PNreadFromText(`
place p (init = 3)
exp t (guard = undefined_variable == 1 && #p >= 0, rate = 1)
`)
	if CompileStats.GuardFallback != 1 || CompileStats.GuardCompiled != 0 {
		t.Errorf("expected the guard to fall back, got %+v", CompileStats)
	}
}
