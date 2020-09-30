package parser

import (
	"fmt"
	"testing"
)

func TestASTExpr1(t *testing.T) {
	x := MakeValue(10.1)
	y := MakeValue(10)
	z := NewASTBinOp("plus", x, y)
	fmt.Println(z)
	result, _ := z.Eval(make(ASTEnv))
	fmt.Println(result)
}

func TestASTExpr2(t *testing.T) {
	x := MakeValue(10.1)
	y := MakeValue(10)
	z := NewASTBinOp("minus", x, y)
	fmt.Println(z)
	env := make(ASTEnv)
	env["test"] = x
	result, _ := z.Eval(env)
	fmt.Println(result)
}

func TestASTExpr3(t *testing.T) {
	x := MakeValue(10.1)
	env := make(ASTEnv)
	env["test"] = x
	x1 := &ASTVar{label: "test"}
	y := MakeValue(10)
	z := NewASTBinOp("minus", x1, y)
	fmt.Println(z)
	result, _ := z.Eval(env)
	fmt.Println(result)
}

func TestASTExpr4(t *testing.T) {
	x := MakeValue(10.1)
	env := make(ASTEnv)
	env["test"] = x
	x1 := &ASTVar{label: "test"}
	y := MakeValue(10)
	z := NewASTBinOp("minus", x1, y)
	env["z"] = z

	lam := func() *ASTValue {
		result, _ := env["z"].(ASTExpr).Eval(env)
		return result
	}

	v := MakeValue(-10000.1)
	env["test"] = v

	result := lam()
	fmt.Println("!!")
	fmt.Println(result)
}

func TestASTExpr5(t *testing.T) {
	x := MakeValue(10.1)
	y := MakeValue(100.1)
	z := MakeValue(1000.1)
	f := NewASTMultiOp("max", x, y, z)
	fmt.Println(f)
	result, _ := f.Eval(make(ASTEnv))
	fmt.Println("test mul max", result)
}

func TestASTExpr6(t *testing.T) {
	x := MakeValue(10.1)
	y := MakeValue(100.1)
	z := MakeValue(1000.1)
	f := NewASTMultiOp("min", x, y, z)
	fmt.Println(f)
	result, _ := f.Eval(make(ASTEnv))
	fmt.Println("test mul min", result)
}

func TestASTExpr7(t *testing.T) {
	x := MakeValue(10.1)
	f := NewASTMultiOp("max", x)
	fmt.Println(f)
	result, _ := f.Eval(make(ASTEnv))
	fmt.Println("test mul max", result)
}

func TestASTExpr8(t *testing.T) {
	f := NewASTMultiOp("max")
	fmt.Println(f)
	result, _ := f.Eval(make(ASTEnv))
	fmt.Println("test mul max", result)
}

func TestASTExpr9(t *testing.T) {
	l := make(ASTList, 0)
	x := MakeValue(1)
	y := MakeValue(2)
	l = append(l, NewASTBinOp("plus", x, y))
	l = append(l, NewASTBinOp("minus", x, y))
	fmt.Println(l)
	result, _ := l.Eval(make(ASTEnv))
	fmt.Println("test list", result)
}

func TestASTExpr10(t *testing.T) {
	env := make(ASTEnv)
	y, res := env["test"].(ASTExpr)
	fmt.Println("test10", y)
	fmt.Println("test10", res)
}
