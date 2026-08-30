package parser

import (
	"math"
	"testing"

	"github.com/okamumu/gospn/pkg/petrinet"
)

// The operators, one case each. These used to print their result and assert nothing.

func TestASTValueBinaryOperators(t *testing.T) {
	tenPointOne, ten := 29.1, 10.0
	for _, tc := range []struct {
		name string
		fn   func(x, y *ASTValue) (*ASTValue, error)
		x, y interface{}
		want interface{}
	}{
		{"plus int+float", plus, ASTInt(10), ASTFloat(29.1), ASTFloat(ten + tenPointOne)},
		{"minus int-float", minus, ASTInt(10), ASTFloat(29.1), ASTFloat(ten - tenPointOne)},
		{"mul", mul, 10, 29.1, ASTFloat(ten * tenPointOne)},
		{"div", div, 10, 29.1, ASTFloat(ten / tenPointOne)},
		{"idiv truncates", idiv, ASTInt(10), ASTInt(29), ASTInt(0)},
		{"eq across int and float", eq, ASTFloat(10), ASTInt(10), ASTBool(true)},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.fn(MakeValue(tc.x), MakeValue(tc.y))
			if err != nil {
				t.Fatal(err)
			}
			if got := res.val; got != tc.want {
				t.Errorf("got %v (%T), want %v (%T)", got, got, tc.want, tc.want)
			}
		})
	}
}

func TestASTValueIfThenElse(t *testing.T) {
	res, err := ite(MakeValue(ASTBool(false)), MakeValue(ASTInt(10)), MakeValue(ASTInt(100)))
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := res.GetInt(); v != ASTInt(100) {
		t.Errorf("ifelse(false, 10, 100) = %v, want 100", v)
	}
}

// An unresolved variable is a string, and an operator on one yields the expression
// written out rather than a number. This is what makes an undefined variable survive
// until something asks for a number -- see petrinet.Net.CheckExpressions.
func TestASTValueUnresolvedOperandStaysSymbolic(t *testing.T) {
	for _, tc := range []struct {
		name string
		x, y interface{}
		want ASTString
	}{
		{"int plus unresolved", ASTInt(10), ASTString("10"), "(10 + 10)"},
		{"unresolved plus unresolved", ASTString("10"), ASTString("10"), "(10 + 10)"},
	} {
		res, err := plus(MakeValue(tc.x), MakeValue(tc.y))
		if err != nil {
			t.Fatal(err)
		}
		v, err := res.GetString()
		if err != nil {
			t.Fatalf("%s: %v", tc.name, err)
		}
		if v != tc.want {
			t.Errorf("%s: got %q, want %q", tc.name, v, tc.want)
		}
	}

	// The condition of an ifelse is no different: it does not become a number.
	res, err := ite(MakeValue(ASTString("x")), MakeValue(ASTInt(10)), MakeValue(ASTInt(100)))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := res.GetInt(); err == nil {
		t.Error("ifelse on an unresolved condition produced an int")
	}
}

func TestASTValueMathFunctions(t *testing.T) {
	for _, tc := range []struct {
		name string
		fn   func(x *ASTValue) (*ASTValue, error)
		x    interface{}
		want ASTFloat
	}{
		{"log(1)", logf, 1, 0},
		{"sqrt(1)", sqrtf, 1.0, 1},
	} {
		res, err := tc.fn(MakeValue(tc.x))
		if err != nil {
			t.Fatalf("%s: %v", tc.name, err)
		}
		if v, _ := res.GetFloat(); v != tc.want {
			t.Errorf("%s = %v, want %v", tc.name, v, tc.want)
		}
	}

	// exp is checked against the same computation rather than a literal.
	res, err := expf(MakeValue(0.1))
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := res.GetFloat(); float64(v) != math.Exp(0.1) {
		t.Errorf("exp(0.1) = %v, want %v", v, math.Exp(0.1))
	}
}

func TestASTValuePowAndMax(t *testing.T) {
	res, err := powf(MakeValue(2), MakeValue(5))
	if err != nil {
		t.Fatal(err)
	}
	// Two ints give an int: pow has an integer path (powi), which is what made a
	// negative exponent loop forever before 0.11.1.
	if v, err := res.GetInt(); err != nil || v != ASTInt(32) {
		t.Errorf("pow(2, 5) = %v (%v), want the int 32", v, err)
	}
	res, err = max(MakeValue(2), MakeValue(5.5))
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := res.GetFloat(); v != ASTFloat(5.5) {
		t.Errorf("max(2, 5.5) = %v, want 5.5", v)
	}
}

// Regression tests for the arithmetic/comparison operators. Unlike the tests
// above these check the returned value, not just that no error occurs.

func TestASTValueMinusInt(t *testing.T) {
	res, err := minus(MakeValue(10), MakeValue(3))
	if err != nil {
		t.Fatal(err)
	}
	v, err := res.GetInt()
	if err != nil {
		t.Fatal(err)
	}
	if v != ASTInt(7) {
		t.Errorf("10 - 3 = %v, want 7", v)
	}
}

func TestASTValueMinusMixed(t *testing.T) {
	res, err := minus(MakeValue(10), MakeValue(2.5))
	if err != nil {
		t.Fatal(err)
	}
	v, err := res.GetFloat()
	if err != nil {
		t.Fatal(err)
	}
	if v != ASTFloat(7.5) {
		t.Errorf("10 - 2.5 = %v, want 7.5", v)
	}
	res, err = minus(MakeValue(2.5), MakeValue(1))
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := res.GetFloat(); v != ASTFloat(1.5) {
		t.Errorf("2.5 - 1 = %v, want 1.5", v)
	}
}

// A reward such as "reward r #P1 - #P2" evaluates as int minus int.
func TestRewardSubtraction(t *testing.T) {
	net, _, err := PNreadFromText(`
		place P1 (init = 5)
		place P2 (init = 2)
		exp T (rate = 1.0)
		arc P1 to T
		arc T to P2
		reward diff #P1 - #P2
	`)
	if err != nil {
		t.Fatal(err)
	}
	rwd, ok := net.GetReward("diff")
	if !ok {
		t.Fatal("reward diff not found")
	}
	if got := rwd([]petrinet.MarkInt{5, 2}); got != 3.0 {
		t.Errorf("#P1 - #P2 with (5,2) = %v, want 3", got)
	}
}

func TestASTValueNeqIntFloat(t *testing.T) {
	res, err := neq(MakeValue(1), MakeValue(1.0))
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := res.GetBool(); v != ASTBool(false) {
		t.Errorf("1 != 1.0 is %v, want false", v)
	}
	res, err = neq(MakeValue(1), MakeValue(2.0))
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := res.GetBool(); v != ASTBool(true) {
		t.Errorf("1 != 2.0 is %v, want true", v)
	}
}

func TestASTValuePowNegativeExponent(t *testing.T) {
	res, err := powf(MakeValue(2), MakeValue(-1))
	if err != nil {
		t.Fatal(err)
	}
	if v, err := res.GetFloat(); err != nil || v != ASTFloat(0.5) {
		t.Errorf("pow(2, -1) = %v (%v), want 0.5", v, err)
	}
}

// Symbolic (unresolved-variable) forms must be well-formed expressions.
func TestASTValueSymbolicParens(t *testing.T) {
	for _, tc := range []struct {
		name string
		fn   func(x, y *ASTValue) (*ASTValue, error)
		want ASTString
	}{
		{"pow", powf, "pow(x, 2.500000e+00)"},
		{"max", max, "max(x, 2.500000e+00)"},
		{"min", min, "min(x, 2.500000e+00)"},
	} {
		res, err := tc.fn(MakeValue("x"), MakeValue(2.5))
		if err != nil {
			t.Fatal(err)
		}
		if v, _ := res.GetString(); v != tc.want {
			t.Errorf("%s: got %q, want %q", tc.name, v, tc.want)
		}
	}
}
