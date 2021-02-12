// Code generated from JSPNL.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // JSPNL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJSPNLListener is a complete listener for a parse tree produced by JSPNLParser.
type BaseJSPNLListener struct{}

var _ JSPNLListener = &BaseJSPNLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJSPNLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJSPNLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJSPNLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJSPNLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseJSPNLListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseJSPNLListener) ExitProg(ctx *ProgContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseJSPNLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseJSPNLListener) ExitStatement(ctx *StatementContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseJSPNLListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseJSPNLListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterNode_declaration is called when production node_declaration is entered.
func (s *BaseJSPNLListener) EnterNode_declaration(ctx *Node_declarationContext) {}

// ExitNode_declaration is called when production node_declaration is exited.
func (s *BaseJSPNLListener) ExitNode_declaration(ctx *Node_declarationContext) {}

// EnterArc_declaration is called when production arc_declaration is entered.
func (s *BaseJSPNLListener) EnterArc_declaration(ctx *Arc_declarationContext) {}

// ExitArc_declaration is called when production arc_declaration is exited.
func (s *BaseJSPNLListener) ExitArc_declaration(ctx *Arc_declarationContext) {}

// EnterReward_declaration is called when production reward_declaration is entered.
func (s *BaseJSPNLListener) EnterReward_declaration(ctx *Reward_declarationContext) {}

// ExitReward_declaration is called when production reward_declaration is exited.
func (s *BaseJSPNLListener) ExitReward_declaration(ctx *Reward_declarationContext) {}

// EnterGroup_declaration is called when production group_declaration is entered.
func (s *BaseJSPNLListener) EnterGroup_declaration(ctx *Group_declarationContext) {}

// ExitGroup_declaration is called when production group_declaration is exited.
func (s *BaseJSPNLListener) ExitGroup_declaration(ctx *Group_declarationContext) {}

// EnterNode_options is called when production node_options is entered.
func (s *BaseJSPNLListener) EnterNode_options(ctx *Node_optionsContext) {}

// ExitNode_options is called when production node_options is exited.
func (s *BaseJSPNLListener) ExitNode_options(ctx *Node_optionsContext) {}

// EnterOption_list is called when production option_list is entered.
func (s *BaseJSPNLListener) EnterOption_list(ctx *Option_listContext) {}

// ExitOption_list is called when production option_list is exited.
func (s *BaseJSPNLListener) ExitOption_list(ctx *Option_listContext) {}

// EnterOption_value is called when production option_value is entered.
func (s *BaseJSPNLListener) EnterOption_value(ctx *Option_valueContext) {}

// ExitOption_value is called when production option_value is exited.
func (s *BaseJSPNLListener) ExitOption_value(ctx *Option_valueContext) {}

// EnterLabel_expression is called when production label_expression is entered.
func (s *BaseJSPNLListener) EnterLabel_expression(ctx *Label_expressionContext) {}

// ExitLabel_expression is called when production label_expression is exited.
func (s *BaseJSPNLListener) ExitLabel_expression(ctx *Label_expressionContext) {}

// EnterUpdate_block is called when production update_block is entered.
func (s *BaseJSPNLListener) EnterUpdate_block(ctx *Update_blockContext) {}

// ExitUpdate_block is called when production update_block is exited.
func (s *BaseJSPNLListener) ExitUpdate_block(ctx *Update_blockContext) {}

// EnterSimple_block is called when production simple_block is entered.
func (s *BaseJSPNLListener) EnterSimple_block(ctx *Simple_blockContext) {}

// ExitSimple_block is called when production simple_block is exited.
func (s *BaseJSPNLListener) ExitSimple_block(ctx *Simple_blockContext) {}

// EnterSimple is called when production simple is entered.
func (s *BaseJSPNLListener) EnterSimple(ctx *SimpleContext) {}

// ExitSimple is called when production simple is exited.
func (s *BaseJSPNLListener) ExitSimple(ctx *SimpleContext) {}

// EnterAssign_expression is called when production assign_expression is entered.
func (s *BaseJSPNLListener) EnterAssign_expression(ctx *Assign_expressionContext) {}

// ExitAssign_expression is called when production assign_expression is exited.
func (s *BaseJSPNLListener) ExitAssign_expression(ctx *Assign_expressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJSPNLListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJSPNLListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFunction_expression is called when production function_expression is entered.
func (s *BaseJSPNLListener) EnterFunction_expression(ctx *Function_expressionContext) {}

// ExitFunction_expression is called when production function_expression is exited.
func (s *BaseJSPNLListener) ExitFunction_expression(ctx *Function_expressionContext) {}

// EnterFunction_args is called when production function_args is entered.
func (s *BaseJSPNLListener) EnterFunction_args(ctx *Function_argsContext) {}

// ExitFunction_args is called when production function_args is exited.
func (s *BaseJSPNLListener) ExitFunction_args(ctx *Function_argsContext) {}

// EnterArgs_list is called when production args_list is entered.
func (s *BaseJSPNLListener) EnterArgs_list(ctx *Args_listContext) {}

// ExitArgs_list is called when production args_list is exited.
func (s *BaseJSPNLListener) ExitArgs_list(ctx *Args_listContext) {}

// EnterArgs_value is called when production args_value is entered.
func (s *BaseJSPNLListener) EnterArgs_value(ctx *Args_valueContext) {}

// ExitArgs_value is called when production args_value is exited.
func (s *BaseJSPNLListener) ExitArgs_value(ctx *Args_valueContext) {}

// EnterNtoken_expression is called when production ntoken_expression is entered.
func (s *BaseJSPNLListener) EnterNtoken_expression(ctx *Ntoken_expressionContext) {}

// ExitNtoken_expression is called when production ntoken_expression is exited.
func (s *BaseJSPNLListener) ExitNtoken_expression(ctx *Ntoken_expressionContext) {}

// EnterEnable_expression is called when production enable_expression is entered.
func (s *BaseJSPNLListener) EnterEnable_expression(ctx *Enable_expressionContext) {}

// ExitEnable_expression is called when production enable_expression is exited.
func (s *BaseJSPNLListener) ExitEnable_expression(ctx *Enable_expressionContext) {}

// EnterLiteral_expression is called when production literal_expression is entered.
func (s *BaseJSPNLListener) EnterLiteral_expression(ctx *Literal_expressionContext) {}

// ExitLiteral_expression is called when production literal_expression is exited.
func (s *BaseJSPNLListener) ExitLiteral_expression(ctx *Literal_expressionContext) {}
