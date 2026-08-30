package parser

import (
	"testing"
)

// These tests used to print their results and assert nothing, which is how the
// arithmetic bugs fixed in 0.11.1 survived. Each one now checks the value.

func evalOK(t *testing.T, e ASTExpr, env ASTEnv) *ASTValue {
	t.Helper()
	res, err := e.Eval(env)
	if err != nil {
		t.Fatalf("%v: %v", e, err)
	}
	return res
}

func wantFloat(t *testing.T, res *ASTValue, want float64, what string) {
	t.Helper()
	v, err := res.GetFloat()
	if err != nil {
		t.Fatalf("%s: %v", what, err)
	}
	if v != ASTFloat(want) {
		t.Errorf("%s = %v, want %v", what, v, want)
	}
}

func TestASTBinOpArithmetic(t *testing.T) {
	// a and b are variables so that the expected values are computed in float64, the
	// way the evaluator computes them. As untyped constants Go would work them out
	// exactly and 10.1 - 10 would be 0.1 rather than 0.09999999999999964.
	a, b := 10.1, 10.0
	x, y := MakeValue(a), MakeValue(10)
	wantFloat(t, evalOK(t, NewASTBinOp("plus", x, y), make(ASTEnv)), a+b, "10.1 + 10")
	wantFloat(t, evalOK(t, NewASTBinOp("minus", x, y), make(ASTEnv)), a-b, "10.1 - 10")
}

func TestASTVarResolvesThroughTheEnv(t *testing.T) {
	a, b := 10.1, 10.0
	env := ASTEnv{"test": MakeValue(a)}
	z := NewASTBinOp("minus", &ASTVar{label: "test"}, MakeValue(10))
	wantFloat(t, evalOK(t, z, env), a-b, "test - 10 with test = 10.1")
}

// Variables are resolved when the expression is evaluated, not when it is built. This
// is what lets a definition file use a parameter before the line that assigns it, and
// it is why an undefined variable is only caught at evaluation time -- see
// petrinet.Net.CheckExpressions.
func TestASTVarIsResolvedLazily(t *testing.T) {
	env := ASTEnv{"test": MakeValue(10.1)}
	env["z"] = NewASTBinOp("minus", &ASTVar{label: "test"}, MakeValue(10))

	a, b := -10000.1, 10.0
	env["test"] = MakeValue(a)
	wantFloat(t, evalOK(t, env["z"].(ASTExpr), env), a-b, "test - 10 after test changed")
}

func TestASTMultiOpMaxMin(t *testing.T) {
	args := []ASTExpr{MakeValue(10.1), MakeValue(100.1), MakeValue(1000.1)}
	wantFloat(t, evalOK(t, NewASTMultiOp("max", args...), make(ASTEnv)), 1000.1, "max(10.1, 100.1, 1000.1)")
	wantFloat(t, evalOK(t, NewASTMultiOp("min", args...), make(ASTEnv)), 10.1, "min(10.1, 100.1, 1000.1)")

	// One argument is its own maximum.
	wantFloat(t, evalOK(t, NewASTMultiOp("max", MakeValue(10.1)), make(ASTEnv)), 10.1, "max(10.1)")
}

// max() with no arguments yields a value holding nothing rather than an error. This
// records what happens; it is not a claim that it is right. The grammar cannot produce
// it -- `max()` does not parse -- so nothing in a definition file reaches this.
func TestASTMultiOpWithNoArgumentsIsEmpty(t *testing.T) {
	res := evalOK(t, NewASTMultiOp("max"), make(ASTEnv))
	if _, err := res.GetFloat(); err == nil {
		t.Error("max() produced a number; it holds nothing")
	}
}

// A list evaluates every element and yields the last one, which is how a block of
// updates ({ #Pn = 1 ... }) is evaluated.
func TestASTListYieldsItsLastValue(t *testing.T) {
	l := ASTList{
		NewASTBinOp("plus", MakeValue(1), MakeValue(2)),
		NewASTBinOp("minus", MakeValue(1), MakeValue(2)),
	}
	res := evalOK(t, l, make(ASTEnv))
	v, err := res.GetInt()
	if err != nil {
		t.Fatal(err)
	}
	if v != ASTInt(-1) {
		t.Errorf("[1+2, 1-2] = %v, want -1 (the last element)", v)
	}
}

func TestASTEnvMissingKeyIsNotAnExpr(t *testing.T) {
	env := make(ASTEnv)
	if e, ok := env["test"].(ASTExpr); ok || e != nil {
		t.Errorf("a missing key gave (%v, %v), want (nil, false)", e, ok)
	}
}
