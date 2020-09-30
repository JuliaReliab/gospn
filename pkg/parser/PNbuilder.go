package parser

import (
	"fmt"
	"strconv"
)

// node
type PNNode struct {
	options map[string]interface{}
}

func newPNNode() *PNNode {
	return &PNNode{
		options: make(map[string]interface{}),
	}
}

type nodeStack []*PNNode

func newNodeStack() nodeStack {
	return make(nodeStack, 0)
}

func (s *nodeStack) push(n *PNNode) {
	*s = append(*s, n)
}

func (s *nodeStack) pop() *PNNode {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

// func (s nodeStack) isempty() bool {
// 	return len(s) == 0
// }

type astStack []ASTExpr

func newAstStack() astStack {
	return make(astStack, 0)
}

func (s *astStack) push(a ASTExpr) {
	*s = append(*s, a)
}

func (s *astStack) pop() ASTExpr {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

// func (s markStack) isempty() bool {
// 	return len(s) == 0
// }

type NetBuilder struct {
	aststack  astStack
	nodestack nodeStack
	labels    []string
	env       ASTEnv
}

func newNetBuilder() *NetBuilder {
	return &NetBuilder{
		aststack:  make(astStack, 0),
		nodestack: make(nodeStack, 0),
		labels:    make([]string, 0),
		env:       make(ASTEnv),
	}
}

func (b *NetBuilder) setNodeOption(label string) {
	node := b.nodestack.pop()
	if label == "" {
		label = fmt.Sprintf("arg%d", len(node.options))
	}
	right := b.aststack.pop()
	node.options[label] = right
	b.nodestack.push(node)
	logger.Print("Set option ", label, " = ", right)
}

func (b *NetBuilder) createNewNode() {
	b.nodestack.push(newPNNode())
	logger.Print("Create a new option node")
}

func (b *NetBuilder) buildNode(ntype string, label string) {
	node := b.nodestack.pop()
	switch ntype {
	case "place":
		node.options["type"] = "place"
		node.options["label"] = label
		logger.Printf("Create a place %s", label)
	case "imm":
		node.options["type"] = "imm"
		node.options["label"] = label
		logger.Printf("Create an imm %s", label)
	case "exp":
		node.options["type"] = "exp"
		node.options["label"] = label
		logger.Printf("Create an exp %s", label)
	case "gen":
		node.options["type"] = "gen"
		node.options["label"] = label
		logger.Printf("Create a gen %s", label)
	case "trans":
		node.options["label"] = label
		logger.Printf("Create a trans %s", label)
	default:
		logger.Panicf("Undified node type %s", ntype)
	}
	b.labels = append(b.labels, label)
	b.env[label] = node
}

func (b *NetBuilder) buildArc(atype string, src string, dest string) {
	node := b.nodestack.pop()
	switch atype {
	case "arc", "iarc", "oarc":
		node.options["type"] = "arc"
		node.options["src"] = src
		node.options["dest"] = dest
		logger.Printf("Create arc %s to %s", src, dest)
	case "harc":
		node.options["type"] = "harc"
		node.options["src"] = src
		node.options["dest"] = dest
		logger.Printf("Create harc %s to %s", src, dest)
	default:
		logger.Panicf("Undified arc type %s", atype)
	}
	b.labels = append(b.labels, src+":"+dest)
	b.env[src+":"+dest] = node
}

func (b *NetBuilder) buildReward(label string) {
	node := b.nodestack.pop()
	node.options["type"] = "reward"
	node.options["label"] = label
	f := b.aststack.pop()
	node.options["formula"] = f
	b.labels = append(b.labels, label)
	b.env[label] = node
	logger.Print("Create reward ", label, " = ", f)
}

func (b *NetBuilder) buildAssignExpression(label string) {
	right := b.aststack.pop()
	b.labels = append(b.labels, label)
	b.env[label] = right
	logger.Print("Put an assign expression ", label, " = ", right)
}

func (b *NetBuilder) buildAssignNTokenExpression() {
	right := b.aststack.pop()
	ntoken := b.aststack.pop()
	var label string
	switch n := ntoken.(type) {
	case *ASTNToken:
		label = n.GetLabel()
	default:
		logger.Panicf("Fail to get NToken as a left value of the assignment")
	}
	a := NewASTAssignNToken(label, right)
	b.aststack.push(a)
	logger.Print("Build an assign ntoken #", label, " = ", right)
}

func (b *NetBuilder) buildValueExpression(label string) {
	tmp, ok := b.env[label]
	if ok {
		switch v := tmp.(type) {
		case ASTExpr:
			b.aststack.push(v)
		default:
			logger.Panic("Value is not ASTExpr ", v)
		}
	} else {
		b.aststack.push(NewASTVar(label))
		logger.Printf("%s was not found. Define the variable %s", label, label)
	}
}

func (b *NetBuilder) buildUnaryExpression(op string) {
	expr := b.aststack.pop()
	var a ASTExpr
	switch op {
	case "+":
		a = NewASTUnaryOp("uplus", expr)
	case "-":
		a = NewASTUnaryOp("uminus", expr)
	case "!":
		a = NewASTUnaryOp("not", expr)
	default:
		logger.Panicf("Undified unary opeartor %s", op)
	}
	b.aststack.push(a)
}

func (b *NetBuilder) buildBinaryExpression(op string) {
	right := b.aststack.pop()
	left := b.aststack.pop()
	var a ASTExpr
	switch op {
	case "+":
		a = NewASTBinOp("plus", left, right)
	case "-":
		a = NewASTBinOp("minus", left, right)
	case "*":
		a = NewASTBinOp("mul", left, right)
	case "div":
		a = NewASTBinOp("idiv", left, right)
	case "/":
		a = NewASTBinOp("div", left, right)
	case "&&":
		a = NewASTBinOp("and", left, right)
	case "||":
		a = NewASTBinOp("or", left, right)
	case "==":
		a = NewASTBinOp("eq", left, right)
	case "!=":
		a = NewASTBinOp("neq", left, right)
	case "<":
		a = NewASTBinOp("lt", left, right)
	case "<=":
		a = NewASTBinOp("lte", left, right)
	case ">":
		a = NewASTBinOp("gt", left, right)
	case ">=":
		a = NewASTBinOp("gte", left, right)
	default:
		logger.Panicf("Undified binary opeartor %s", op)
	}
	b.aststack.push(a)
}

func (b *NetBuilder) buildIfThenElseExpression(_ string) {
	third := b.aststack.pop()
	second := b.aststack.pop()
	first := b.aststack.pop()
	a := NewASTTriOp("ite", first, second, third)
	b.aststack.push(a)
}

func (b *NetBuilder) buildNToken(label string) {
	a := NewASTNToken(label)
	b.aststack.push(a)
}

func (b *NetBuilder) buildEnable(label string) {
	a := NewASTEnableCond(label)
	b.aststack.push(a)
}

func (b *NetBuilder) buildIntegerLiteral(value string) {
	v, err := strconv.Atoi(value)
	if err != nil {
		logger.Panicf("Fail to convert %s to int", value)
	}
	a := MakeValue(v)
	b.aststack.push(a)
}

func (b *NetBuilder) buildDoubleLiteral(value string) {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		logger.Panicf("Fail to convert %s to float", value)
	}
	a := MakeValue(v)
	b.aststack.push(a)
}

func (b *NetBuilder) buildBooleanLiteral(value string) {
	v, err := strconv.ParseBool(value)
	if err != nil {
		logger.Panicf("Fail to convert %s to boolean", value)
	}
	a := MakeValue(v)
	b.aststack.push(a)
}

func (b *NetBuilder) buildStringLiteral(value string) {
	a := MakeValue(value)
	b.aststack.push(a)
}

func (b *NetBuilder) buildFunc(funcname string) {
	var a ASTExpr
	switch funcname {
	case "exp":
		a = buildExpFunc(b.nodestack.pop())
	case "log":
		a = buildLogFunc(b.nodestack.pop())
	case "sqrt":
		a = buildSqrtFunc(b.nodestack.pop())
	case "pow":
		a = buildPowFunc(b.nodestack.pop())
	case "min":
		a = buildMinFunc(b.nodestack.pop())
	case "max":
		a = buildMaxFunc(b.nodestack.pop())
	case "det":
		a = buildConstDist(b.nodestack.pop())
	case "unif":
		a = buildUnifDist(b.nodestack.pop())
	case "expdist":
		a = buildExpDist(b.nodestack.pop())
	default:
		logger.Panicf("Function %s is not defined yet", funcname)
	}
	b.aststack.push(a)
}

func buildExpFunc(x *PNNode) ASTExpr {
	args := make(ASTList, 1, 1)
	args[0] = MakeValue(0.0)
	for key, value := range x.options {
		switch key {
		case "x", "arg0":
			args[0] = value.(ASTExpr)
		default:
			logger.Printf("Ignore the arg %s in exp", key)
		}
	}
	return NewASTUnaryOp("expf", args[0])
}

func buildLogFunc(x *PNNode) ASTExpr {
	args := make(ASTList, 1, 1)
	args[0] = MakeValue(1.0)
	for key, value := range x.options {
		switch key {
		case "x", "arg0":
			args[0] = value.(ASTExpr)
		default:
			logger.Printf("Ignore the arg %s in log", key)
		}
	}
	return NewASTUnaryOp("logf", args[0])
}

func buildSqrtFunc(x *PNNode) ASTExpr {
	args := make(ASTList, 1, 1)
	args[0] = MakeValue(1.0)
	for key, value := range x.options {
		switch key {
		case "x", "arg0":
			args[0] = value.(ASTExpr)
		default:
			logger.Printf("Ignore the arg %s in sqrt", key)
		}
	}
	return NewASTUnaryOp("sqrtf", args[0])
}

func buildPowFunc(x *PNNode) ASTExpr {
	args := make(ASTList, 2, 2)
	args[0] = MakeValue(1)
	args[1] = MakeValue(1)
	for key, value := range x.options {
		switch key {
		case "x", "arg0":
			args[0] = value.(ASTExpr)
		case "n", "arg1":
			args[1] = value.(ASTExpr)
		}
	}
	return NewASTBinOp("powf", args[0], args[1])
}

func buildMaxFunc(x *PNNode) ASTExpr {
	args := make(ASTList, 0)
	for _, value := range x.options {
		args = append(args, value.(ASTExpr))
	}
	return NewASTMultiOp("max", args...)
}

func buildMinFunc(x *PNNode) ASTExpr {
	args := make(ASTList, 0)
	for _, value := range x.options {
		args = append(args, value.(ASTExpr))
	}
	return NewASTMultiOp("min", args...)
}

func buildConstDist(x *PNNode) ASTExpr {
	args := make(ASTList, 1, 1)
	args[0] = MakeValue(1.0)
	for key, value := range x.options {
		switch key {
		case "value", "arg0":
			args[0] = value.(ASTExpr)
		}
	}
	return NewASTUnaryOp("constdist", args[0])
}

func buildUnifDist(x *PNNode) ASTExpr {
	args := make(ASTList, 2, 2)
	args[0] = MakeValue(0.0)
	args[1] = MakeValue(1.0)
	for key, value := range x.options {
		switch key {
		case "min", "arg0":
			args[0] = value.(ASTExpr)
		case "max", "arg1":
			args[1] = value.(ASTExpr)
		}
	}
	return NewASTBinOp("unifdist", args[0], args[1])
}

func buildExpDist(x *PNNode) ASTExpr {
	args := make(ASTList, 1, 1)
	args[0] = MakeValue(1.0)
	for key, value := range x.options {
		switch key {
		case "rate", "arg0":
			args[0] = value.(ASTExpr)
		}
	}
	return NewASTUnaryOp("expdist", args[0])
}

func (b *NetBuilder) setUpdateBlockEnd() {
	a := NewASTNop()
	b.aststack.push(a)
	logger.Printf("Set an update block end")
}

func isblockend(x ASTExpr) bool {
	switch x.(type) {
	case *ASTNop:
		return true
	default:
		return false
	}
}

func (b *NetBuilder) buildUpdateBlock() {
	list := make(ASTList, 0)
	a := b.aststack.pop()
	for isblockend(a) == false {
		list = append(list, a)
		a = b.aststack.pop()
	}
	node := b.nodestack.pop()
	node.options["update"] = list
	b.nodestack.push(node)
	logger.Printf("Put an update block")
}

func (b *NetBuilder) parserError() {
	logger.Panic("Parser error. Stop to run")
}
