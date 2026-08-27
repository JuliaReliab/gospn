package parser

import (
	"fmt"
	"testing"

	"github.com/okamumu/gospn/pkg/petrinet"
)

func TestASTValue1(t *testing.T) {
	a := ASTInt(10)
	b := ASTFloat(29.1)
	x := MakeValue(a)
	y := MakeValue(b)
	z, _ := plus(x, y)
	fmt.Println(z)
}

func TestASTValue2(t *testing.T) {
	a := ASTInt(10)
	b := ASTFloat(29.1)
	x := MakeValue(a)
	y := MakeValue(b)
	z, _ := minus(x, y)
	fmt.Println(z)
}

func TestASTValue3(t *testing.T) {
	x := MakeValue(10)
	y := MakeValue(29.1)
	z, _ := mul(x, y)
	fmt.Println(z)
}

func TestASTValue4(t *testing.T) {
	a := ASTInt(10)
	b := ASTFloat(29.1)
	x := MakeValue(a)
	y := MakeValue(b)
	z, _ := div(x, y)
	fmt.Println(z)
}

func TestASTValue5(t *testing.T) {
	a := ASTInt(10)
	b := ASTInt(29)
	x := MakeValue(a)
	y := MakeValue(b)
	z, _ := idiv(x, y)
	fmt.Println(z)
}

func TestASTValue6(t *testing.T) {
	a := ASTFloat(10)
	b := ASTInt(10)
	x := MakeValue(a)
	y := MakeValue(b)
	z, _ := eq(x, y)
	fmt.Println(z)
}

func TestASTValue7(t *testing.T) {
	a := ASTBool(false)
	b := ASTInt(10)
	c := ASTInt(100)
	x := MakeValue(a)
	y := MakeValue(b)
	z := MakeValue(c)
	v, _ := ite(x, y, z)
	fmt.Println(v)
}

func TestASTValue8(t *testing.T) {
	a := ASTInt(10)
	b := ASTString("10")
	x := MakeValue(a)
	y := MakeValue(b)
	z, _ := plus(x, y)
	fmt.Println(z)
}

func TestASTValue9(t *testing.T) {
	a := ASTString("10")
	b := ASTString("10")
	x := MakeValue(a)
	y := MakeValue(b)
	z, _ := plus(x, y)
	fmt.Println(z)
}

func TestASTValue10(t *testing.T) {
	a := ASTString("x")
	b := ASTInt(10)
	c := ASTInt(100)
	x := MakeValue(a)
	y := MakeValue(b)
	z := MakeValue(c)
	v, _ := ite(x, y, z)
	res, err := v.GetInt()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestASTValue11(t *testing.T) {
	a := MakeValue(0.1)
	res, err := expf(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestASTValue12(t *testing.T) {
	a := MakeValue(1)
	res, err := logf(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestASTValue13(t *testing.T) {
	a := MakeValue(1.0)
	res, err := sqrtf(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestASTValue14(t *testing.T) {
	a := MakeValue(2)
	b := MakeValue(5)
	res, err := powf(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestASTValue15(t *testing.T) {
	a := MakeValue(2)
	b := MakeValue(5.5)
	res, err := max(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
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
	net, _ := PNreadFromText(`
		place P1 (init = 5)
		place P2 (init = 2)
		exp T (rate = 1.0)
		arc P1 to T
		arc T to P2
		reward diff #P1 - #P2
	`)
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
