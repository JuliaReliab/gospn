// Code generated from JSPNL.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // JSPNL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JSPNLListener is a complete listener for a parse tree produced by JSPNLParser.
type JSPNLListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterNode_declaration is called when entering the node_declaration production.
	EnterNode_declaration(c *Node_declarationContext)

	// EnterArc_declaration is called when entering the arc_declaration production.
	EnterArc_declaration(c *Arc_declarationContext)

	// EnterReward_declaration is called when entering the reward_declaration production.
	EnterReward_declaration(c *Reward_declarationContext)

	// EnterNode_options is called when entering the node_options production.
	EnterNode_options(c *Node_optionsContext)

	// EnterOption_list is called when entering the option_list production.
	EnterOption_list(c *Option_listContext)

	// EnterOption_value is called when entering the option_value production.
	EnterOption_value(c *Option_valueContext)

	// EnterLabel_expression is called when entering the label_expression production.
	EnterLabel_expression(c *Label_expressionContext)

	// EnterUpdate_block is called when entering the update_block production.
	EnterUpdate_block(c *Update_blockContext)

	// EnterSimple_block is called when entering the simple_block production.
	EnterSimple_block(c *Simple_blockContext)

	// EnterSimple is called when entering the simple production.
	EnterSimple(c *SimpleContext)

	// EnterAssign_expression is called when entering the assign_expression production.
	EnterAssign_expression(c *Assign_expressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFunction_expression is called when entering the function_expression production.
	EnterFunction_expression(c *Function_expressionContext)

	// EnterFunction_args is called when entering the function_args production.
	EnterFunction_args(c *Function_argsContext)

	// EnterArgs_list is called when entering the args_list production.
	EnterArgs_list(c *Args_listContext)

	// EnterArgs_value is called when entering the args_value production.
	EnterArgs_value(c *Args_valueContext)

	// EnterNtoken_expression is called when entering the ntoken_expression production.
	EnterNtoken_expression(c *Ntoken_expressionContext)

	// EnterEnable_expression is called when entering the enable_expression production.
	EnterEnable_expression(c *Enable_expressionContext)

	// EnterLiteral_expression is called when entering the literal_expression production.
	EnterLiteral_expression(c *Literal_expressionContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitNode_declaration is called when exiting the node_declaration production.
	ExitNode_declaration(c *Node_declarationContext)

	// ExitArc_declaration is called when exiting the arc_declaration production.
	ExitArc_declaration(c *Arc_declarationContext)

	// ExitReward_declaration is called when exiting the reward_declaration production.
	ExitReward_declaration(c *Reward_declarationContext)

	// ExitNode_options is called when exiting the node_options production.
	ExitNode_options(c *Node_optionsContext)

	// ExitOption_list is called when exiting the option_list production.
	ExitOption_list(c *Option_listContext)

	// ExitOption_value is called when exiting the option_value production.
	ExitOption_value(c *Option_valueContext)

	// ExitLabel_expression is called when exiting the label_expression production.
	ExitLabel_expression(c *Label_expressionContext)

	// ExitUpdate_block is called when exiting the update_block production.
	ExitUpdate_block(c *Update_blockContext)

	// ExitSimple_block is called when exiting the simple_block production.
	ExitSimple_block(c *Simple_blockContext)

	// ExitSimple is called when exiting the simple production.
	ExitSimple(c *SimpleContext)

	// ExitAssign_expression is called when exiting the assign_expression production.
	ExitAssign_expression(c *Assign_expressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFunction_expression is called when exiting the function_expression production.
	ExitFunction_expression(c *Function_expressionContext)

	// ExitFunction_args is called when exiting the function_args production.
	ExitFunction_args(c *Function_argsContext)

	// ExitArgs_list is called when exiting the args_list production.
	ExitArgs_list(c *Args_listContext)

	// ExitArgs_value is called when exiting the args_value production.
	ExitArgs_value(c *Args_valueContext)

	// ExitNtoken_expression is called when exiting the ntoken_expression production.
	ExitNtoken_expression(c *Ntoken_expressionContext)

	// ExitEnable_expression is called when exiting the enable_expression production.
	ExitEnable_expression(c *Enable_expressionContext)

	// ExitLiteral_expression is called when exiting the literal_expression production.
	ExitLiteral_expression(c *Literal_expressionContext)
}
