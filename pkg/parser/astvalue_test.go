package parser

import (
	"fmt"
	"testing"
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
