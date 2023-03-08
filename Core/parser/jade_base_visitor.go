// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Jade

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseJadeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJadeVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitIterationStmt(ctx *IterationStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitFunctionStmt(ctx *FunctionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitAllStmts(ctx *AllStmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitImportStmt(ctx *ImportStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitUndefineStmt(ctx *UndefineStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitInheritList(ctx *InheritListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitClassdef(ctx *ClassdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitVarAssign(ctx *VarAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitFuncAssign(ctx *FuncAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitGetAttributes(ctx *GetAttributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitFuncExpr(ctx *FuncExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJadeVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}
