// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Jade

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by JadeParser.
type JadeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JadeParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by JadeParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by JadeParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by JadeParser#iterationStmt.
	VisitIterationStmt(ctx *IterationStmtContext) interface{}

	// Visit a parse tree produced by JadeParser#functionStmt.
	VisitFunctionStmt(ctx *FunctionStmtContext) interface{}

	// Visit a parse tree produced by JadeParser#allStmts.
	VisitAllStmts(ctx *AllStmtsContext) interface{}

	// Visit a parse tree produced by JadeParser#importStmt.
	VisitImportStmt(ctx *ImportStmtContext) interface{}

	// Visit a parse tree produced by JadeParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by JadeParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by JadeParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by JadeParser#undefineStmt.
	VisitUndefineStmt(ctx *UndefineStmtContext) interface{}

	// Visit a parse tree produced by JadeParser#inheritList.
	VisitInheritList(ctx *InheritListContext) interface{}

	// Visit a parse tree produced by JadeParser#classdef.
	VisitClassdef(ctx *ClassdefContext) interface{}

	// Visit a parse tree produced by JadeParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by JadeParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by JadeParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by JadeParser#varAssign.
	VisitVarAssign(ctx *VarAssignContext) interface{}

	// Visit a parse tree produced by JadeParser#funcAssign.
	VisitFuncAssign(ctx *FuncAssignContext) interface{}

	// Visit a parse tree produced by JadeParser#getAttributes.
	VisitGetAttributes(ctx *GetAttributesContext) interface{}

	// Visit a parse tree produced by JadeParser#funcExpr.
	VisitFuncExpr(ctx *FuncExprContext) interface{}

	// Visit a parse tree produced by JadeParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by JadeParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by JadeParser#atom.
	VisitAtom(ctx *AtomContext) interface{}
}
