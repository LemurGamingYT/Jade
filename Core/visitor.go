package Core

import (
	jade "Jade/Core/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Visitor struct {
	jade.BaseJadeVisitor

	env Env
}

func (v *Visitor) Visit(tree antlr.ParseTree) any {
	switch val := tree.(type) {
	case *jade.ParseContext:
		return v.VisitParse(val)
	case *jade.BlockContext:
		return v.VisitBlock(val)
	case *jade.StmtContext:
		return v.VisitStmt(val)
	case *jade.CallContext:
		return v.VisitCall(val)
	case *jade.VarAssignContext:
		return v.VisitVarAssign(val)
	case *jade.FuncAssignContext:
		return v.VisitFuncAssign(val)
	case *jade.IterationStmtContext:
		return v.VisitIterationStmt(val)
	// Add more cases for each context type defined in your Jade grammar
	default:
		panic(fmt.Sprintf("unsupported type: %T", tree))
	}

	return nil
}

func (v *Visitor) VisitCall(ctx *jade.CallContext) interface{} {
	fmt.Println("Call")

	name := ctx.ID().GetText()
	fmt.Println(name)

	return name
}

func (v *Visitor) VisitExpr(ctx *jade.ExprContext) interface{} {
	fmt.Println("Expr")
	if ctx.Atom() != nil {
		return v.VisitAtom(ctx.Atom().(*jade.AtomContext))
	} else if ctx.Call() != nil {
		return v.VisitCall(ctx.Call().(*jade.CallContext))
	}

	return nil
}

func (v *Visitor) VisitAtom(ctx *jade.AtomContext) interface{} {
	fmt.Println("Atom")

	text := ctx.GetText()
	switch {
	case ctx.STRING() != nil:
		return NewStringObj(text)

	case ctx.INT() != nil:
		return NewIntObj(NewBigInt(text))

	case ctx.FLOAT() != nil:
		return NewFloatObj(NewBigFloat(text))

	case ctx.BOOL() != nil:
		return NewBoolObj(NewBool(text))

	default:
		return NewNullObj()
	}
}

func NewVisitor(env Env) *Visitor {
	return &Visitor{env: env}
}
