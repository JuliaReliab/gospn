package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type PNListener struct {
	*BaseJSPNLListener
	builder *NetBuilder
}

func NewPNListener() *PNListener {
	return &PNListener{
		BaseJSPNLListener: new(BaseJSPNLListener),
		builder:           newNetBuilder(),
	}
}

func (l *PNListener) VisitErrorNode(node antlr.ErrorNode) {
	l.builder.parserError()
}

func (l *PNListener) EnterDeclaration(c *DeclarationContext) {
	l.builder.createNewNode()
}

func (l *PNListener) ExitNode_declaration(c *Node_declarationContext) {
	l.builder.buildNode(c.node.GetText(), c.id.GetText())
}

func (l *PNListener) ExitArc_declaration(c *Arc_declarationContext) {
	l.builder.buildArc(c.arctype.GetText(), c.srcName.GetText(), c.destName.GetText())
}

func (l *PNListener) ExitReward_declaration(c *Reward_declarationContext) {
	l.builder.buildReward(c.id.GetText())
}

func (l *PNListener) ExitGroup_declaration(c *Group_declarationContext) {
	l.builder.buildGroup(c.id.GetText())
}

func (l *PNListener) ExitLabel_expression(c *Label_expressionContext) {
	l.builder.setNodeOption(c.id.GetText())
}

func (l *PNListener) EnterUpdate_block(c *Update_blockContext) {
	l.builder.setUpdateBlockEnd()
}

func (l *PNListener) ExitUpdate_block(c *Update_blockContext) {
	l.builder.buildUpdateBlock()
}

func (l *PNListener) ExitAssign_expression(c *Assign_expressionContext) {
	switch c.exprtype {
	case 1:
		l.builder.buildAssignExpression(c.id.GetText())
	case 2:
		l.builder.buildAssignNTokenExpression()
	default:
	}
}

func (l *PNListener) ExitExpression(c *ExpressionContext) {
	switch c.nodetype {
	case 1:
		l.builder.buildUnaryExpression(c.op.GetText())
	case 2, 3, 4, 5, 6, 7:
		l.builder.buildBinaryExpression(c.op.GetText())
	case 8:
		l.builder.buildIfThenElseExpression(c.op.GetText())
	case 9, 10, 11:
		// nop
	case 12:
		l.builder.buildValueExpression(c.id.GetText())
	case 13, 14:
		// nop
	default:
	}
}

func (l *PNListener) EnterFunction_expression(c *Function_expressionContext) {
	l.builder.createNewNode()
}

func (l *PNListener) ExitFunction_expression(ctx *Function_expressionContext) {
	l.builder.buildFunc(ctx.id.GetText())
}

func (l *PNListener) ExitArgs_value(c *Args_valueContext) {
	l.builder.setNodeOption("")
}

func (l *PNListener) ExitNtoken_expression(c *Ntoken_expressionContext) {
	l.builder.buildNToken(c.id.GetText())
}

func (l *PNListener) ExitEnable_expression(c *Enable_expressionContext) {
	l.builder.buildEnable(c.id.GetText())
}

func (l *PNListener) ExitLiteral_expression(c *Literal_expressionContext) {
	switch c.littype {
	case 1:
		l.builder.buildIntegerLiteral(c.val.GetText())
	case 2:
		l.builder.buildDoubleLiteral(c.val.GetText())
	case 3:
		l.builder.buildBooleanLiteral(c.val.GetText())
	case 4:
		l.builder.buildStringLiteral(c.val.GetText())
	default:
	}
}
