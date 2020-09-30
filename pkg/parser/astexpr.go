package parser

import (
	"../petrinet"
	"errors"
	"fmt"
)

type ASTExpr interface {
	Eval(env ASTEnv) (*ASTValue, error)
	EvalWithMark(net *petrinet.Net, mark []petrinet.MarkInt, env ASTEnv) (*ASTValue, error)
}

// The type for ASTEnv. This is the environment to store the assinged values.
type ASTEnv map[string]interface{}

// The Eval function for ASTValue. *ASTValue is one of ASTExpr.
// The definition of ASTValue is in astvalue.go
func (a *ASTValue) Eval(env ASTEnv) (*ASTValue, error) {
	return a, nil
}

func (a *ASTValue) EvalWithMark(net *petrinet.Net, mark []petrinet.MarkInt, env ASTEnv) (*ASTValue, error) {
	return a, nil
}

// The variable for ASTExpr
type ASTVar struct {
	label string
}

// The function to create
func NewASTVar(label string) *ASTVar {
	return &ASTVar{
		label: label,
	}
}

// The Eval function for ASTVar. If the variable is not defined in ASTEnv,
// it returns a string of the label.
func (a *ASTVar) Eval(env ASTEnv) (*ASTValue, error) {
	value, ok := env[a.label]
	if ok == false {
		return MakeValue(a.label), nil
	}
	expr, ok := value.(ASTExpr)
	if ok == false {
		return MakeValue(nil), errors.New("var is not expr")
	}
	return expr.Eval(env)
}

func (a *ASTVar) EvalWithMark(net *petrinet.Net, mark []petrinet.MarkInt, env ASTEnv) (*ASTValue, error) {
	value, ok := env[a.label]
	if ok == false {
		return MakeValue(a.label), nil
	}
	expr, ok := value.(ASTExpr)
	if ok == false {
		return MakeValue(nil), errors.New("var is not expr")
	}
	return expr.EvalWithMark(net, mark, env)
}

// The structure of unary operator.
// Operators: uplus, uminus, not, expf, logf, sqrtf
type ASTUnaryOp struct {
	op   string
	elem ASTExpr
}

func NewASTUnaryOp(op string, elem ASTExpr) ASTExpr {
	return &ASTUnaryOp{
		op:   op,
		elem: elem,
	}
}

func (a *ASTUnaryOp) Eval(env ASTEnv) (*ASTValue, error) {
	val, err := a.elem.Eval(env)
	if err != nil {
		return val, err
	}
	switch a.op {
	case "uplus":
		return uplus(val)
	case "uminus":
		return uminus(val)
	case "not":
		return not(val)
	case "expf":
		return expf(val)
	case "logf":
		return logf(val)
	case "sqrtf":
		return sqrtf(val)
	case "constdist":
		dist := ASTDist{
			name:   MakeValue("constant"),
			params: []*ASTValue{val},
		}
		return MakeValue(dist), nil
	case "expdist":
		dist := ASTDist{
			name:   MakeValue("exponential"),
			params: []*ASTValue{val},
		}
		return MakeValue(dist), nil
	default:
		return MakeValue(nil), errors.New(fmt.Sprintf("unarry operation: %s is undefined", a.op))
	}
}

func (a *ASTUnaryOp) EvalWithMark(net *petrinet.Net, mark []petrinet.MarkInt, env ASTEnv) (*ASTValue, error) {
	val, err := a.elem.EvalWithMark(net, mark, env)
	if err != nil {
		return val, err
	}
	switch a.op {
	case "uplus":
		return uplus(val)
	case "uminus":
		return uminus(val)
	case "not":
		return not(val)
	case "expf":
		return expf(val)
	case "logf":
		return logf(val)
	case "sqrtf":
		return sqrtf(val)
	case "constdist":
		dist := ASTDist{
			name:   MakeValue("constant"),
			params: []*ASTValue{val},
		}
		return MakeValue(dist), nil
	case "expdist":
		dist := ASTDist{
			name:   MakeValue("exponential"),
			params: []*ASTValue{val},
		}
		return MakeValue(dist), nil
	default:
		return MakeValue(nil), errors.New(fmt.Sprintf("unarry operation: %s is undefined", a.op))
	}
}

// The structure of binary operator.
// operators: plus, minus, mul, idiv, div
// operators: and, or
// operators: eq, neq, lt, lte, gt, gte
// operators: powf, max, min
// TODO: mod operator should be implemented
type ASTBinOp struct {
	op    string
	left  ASTExpr
	right ASTExpr
}

func NewASTBinOp(op string, left ASTExpr, right ASTExpr) ASTExpr {
	return &ASTBinOp{
		op:    op,
		left:  left,
		right: right,
	}
}

// The function for multi args.
// This is only available for plus, mul, and, or, max, min
func NewASTMultiOp(op string, args ...ASTExpr) ASTExpr {
	if len(args) == 0 {
		return MakeValue(nil)
	}
	x := args[0]
	for _, v := range args[1:] {
		x = NewASTBinOp(op, x, v)
	}
	return x
}

func (a *ASTBinOp) Eval(env ASTEnv) (*ASTValue, error) {
	val1, err1 := a.left.Eval(env)
	if err1 != nil {
		return val1, err1
	}
	val2, err2 := a.right.Eval(env)
	if err2 != nil {
		return val2, err2
	}
	switch a.op {
	case "plus":
		return plus(val1, val2)
	case "minus":
		return minus(val1, val2)
	case "mul":
		return mul(val1, val2)
	case "idiv":
		return idiv(val1, val2)
	case "div":
		return div(val1, val2)
	case "and":
		return and(val1, val2)
	case "or":
		return or(val1, val2)
	case "eq":
		return eq(val1, val2)
	case "neq":
		return neq(val1, val2)
	case "lt":
		return lt(val1, val2)
	case "lte":
		return lte(val1, val2)
	case "gt":
		return gt(val1, val2)
	case "gte":
		return gte(val1, val2)
	case "powf":
		return powf(val1, val2)
	case "max":
		return max(val1, val2)
	case "min":
		return min(val1, val2)
	case "unifdist":
		dist := ASTDist{
			name:   MakeValue("uniform"),
			params: []*ASTValue{val1, val2},
		}
		return MakeValue(dist), nil
	default:
		return MakeValue(nil), errors.New(fmt.Sprintf("binary operation: %s is undefined", a.op))
	}
}

func (a *ASTBinOp) EvalWithMark(net *petrinet.Net, mark []petrinet.MarkInt, env ASTEnv) (*ASTValue, error) {
	val1, err1 := a.left.EvalWithMark(net, mark, env)
	if err1 != nil {
		return val1, err1
	}
	val2, err2 := a.right.EvalWithMark(net, mark, env)
	if err2 != nil {
		return val2, err2
	}
	switch a.op {
	case "plus":
		return plus(val1, val2)
	case "minus":
		return minus(val1, val2)
	case "mul":
		return mul(val1, val2)
	case "idiv":
		return idiv(val1, val2)
	case "div":
		return div(val1, val2)
	case "and":
		return and(val1, val2)
	case "or":
		return or(val1, val2)
	case "eq":
		return eq(val1, val2)
	case "neq":
		return neq(val1, val2)
	case "lt":
		return lt(val1, val2)
	case "lte":
		return lte(val1, val2)
	case "gt":
		return gt(val1, val2)
	case "gte":
		return gte(val1, val2)
	case "powf":
		return powf(val1, val2)
	case "max":
		return max(val1, val2)
	case "min":
		return min(val1, val2)
	case "det":
		return min(val1, val2)
	case "unifdist":
		dist := ASTDist{
			name:   MakeValue("uniform"),
			params: []*ASTValue{val1, val2},
		}
		return MakeValue(dist), nil
	default:
		return MakeValue(nil), errors.New(fmt.Sprintf("binary operation: %s is undefined", a.op))
	}
}

// The structure of trinary operator
// operators: ite (if-then-else)
type ASTTriOp struct {
	op     string
	first  ASTExpr
	second ASTExpr
	third  ASTExpr
}

func NewASTTriOp(op string, first ASTExpr, second ASTExpr, third ASTExpr) ASTExpr {
	return &ASTTriOp{
		op:     op,
		first:  first,
		second: second,
		third:  third,
	}
}

func (a *ASTTriOp) Eval(env ASTEnv) (*ASTValue, error) {
	val1, err1 := a.first.Eval(env)
	if err1 != nil {
		return val1, err1
	}
	val2, err2 := a.second.Eval(env)
	if err2 != nil {
		return val2, err2
	}
	val3, err3 := a.third.Eval(env)
	if err3 != nil {
		return val3, err3
	}
	switch a.op {
	case "ite":
		return ite(val1, val2, val3)
	default:
		return MakeValue(nil), errors.New(fmt.Sprintf("trinary operation: %s is undefined", a.op))
	}
}

func (a *ASTTriOp) EvalWithMark(net *petrinet.Net, mark []petrinet.MarkInt, env ASTEnv) (*ASTValue, error) {
	val1, err1 := a.first.EvalWithMark(net, mark, env)
	if err1 != nil {
		return val1, err1
	}
	val2, err2 := a.second.EvalWithMark(net, mark, env)
	if err2 != nil {
		return val2, err2
	}
	val3, err3 := a.third.EvalWithMark(net, mark, env)
	if err3 != nil {
		return val3, err3
	}
	switch a.op {
	case "ite":
		return ite(val1, val2, val3)
	default:
		return MakeValue(nil), errors.New(fmt.Sprintf("trinary operation: %s is undefined", a.op))
	}
}

///

type ASTNToken struct {
	label string
}

func NewASTNToken(label string) ASTExpr {
	return &ASTNToken{
		label: label,
	}
}

func (a *ASTNToken) GetLabel() string {
	return a.label
}

func (a *ASTNToken) Eval(env ASTEnv) (*ASTValue, error) {
	return MakeValue("#" + a.label), nil
}

func (a *ASTNToken) EvalWithMark(net *petrinet.Net, mark []petrinet.MarkInt, env ASTEnv) (*ASTValue, error) {
	place, ok := net.GetPlace(a.label)
	if ok {
		return MakeValue(ASTInt(mark[place.GetIndex()])), nil
	} else {
		return MakeValue(nil), errors.New(fmt.Sprintf("Fail to find the place %s in NToken", a.label))
	}
}

///

type ASTEnableCond struct {
	label string
}

func NewASTEnableCond(label string) ASTExpr {
	return &ASTEnableCond{
		label: label,
	}
}

func (a *ASTEnableCond) Eval(env ASTEnv) (*ASTValue, error) {
	return MakeValue("?" + a.label), nil
}

func (a *ASTEnableCond) EvalWithMark(net *petrinet.Net, mark []petrinet.MarkInt, env ASTEnv) (*ASTValue, error) {
	tr, ok := net.GetTrans(a.label)
	if ok {
		return MakeValue(tr.IsEnabled(net, mark) == petrinet.ENABLE), nil
	} else {
		return MakeValue(nil), errors.New(fmt.Sprintf("Fail to get trans %sin EnableCond", a.label))
	}
}

///

type ASTList []ASTExpr

func (a ASTList) Eval(env ASTEnv) (*ASTValue, error) {
	var finalresult *ASTValue
	for _, x := range a {
		result, err := x.Eval(env)
		if err != nil {
			return result, err
		}
		finalresult = result
	}
	return finalresult, nil
}

func (a ASTList) EvalWithMark(net *petrinet.Net, mark []petrinet.MarkInt, env ASTEnv) (*ASTValue, error) {
	var finalresult *ASTValue
	for _, x := range a {
		result, err := x.EvalWithMark(net, mark, env)
		if err != nil {
			return result, err
		}
		finalresult = result
	}
	return finalresult, nil
}

///

type ASTAssignNToken struct {
	label string
	right ASTExpr
}

func NewASTAssignNToken(label string, right ASTExpr) ASTExpr {
	return &ASTAssignNToken{
		label: label,
		right: right,
	}
}

func (a *ASTAssignNToken) Eval(env ASTEnv) (*ASTValue, error) {
	result, err := a.right.Eval(env)
	if err != nil {
		return result, err
	}
	switch v := result.val.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("#%s = %d", a.label, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("#%s = %e", a.label, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("#%s = %s", a.label, v)), nil
	default:
		return MakeValue(nil), errors.New(fmt.Sprintf("The type of right-hand term in AssignNToken is undefined."))
	}
}

func (a *ASTAssignNToken) EvalWithMark(net *petrinet.Net, mark []petrinet.MarkInt, env ASTEnv) (*ASTValue, error) {
	result, err := a.right.EvalWithMark(net, mark, env)
	if err != nil {
		return result, err
	}
	switch v := result.val.(type) {
	case ASTInt:
		if place, ok := net.GetPlace(a.label); ok {
			mark[place.GetIndex()] = petrinet.MarkInt(v)
			return MakeValue(v), nil
		} else {
			return MakeValue(nil), errors.New(fmt.Sprintf("Fail to get place %s in AssignNToken", a.label))
		}
	case ASTFloat:
		return MakeValue(nil), errors.New("The right-hand value in AssignNToken should be int")
	case ASTString:
		return MakeValue(nil), errors.New("The right-hand value in AssignNToken should be int")
	default:
		return MakeValue(nil), errors.New(fmt.Sprintf("The type of right-hand term in AssignNToken is undefined."))
	}
}

///

type ASTNop struct{}

func NewASTNop() ASTExpr {
	return &ASTNop{}
}

func (a *ASTNop) Eval(env ASTEnv) (*ASTValue, error) {
	return nil, nil
}

func (a *ASTNop) EvalWithMark(net *petrinet.Net, mark []petrinet.MarkInt, env ASTEnv) (*ASTValue, error) {
	return nil, nil
}
