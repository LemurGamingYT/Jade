package Core

import jade "Jade/Core/parser"

type FuncObj struct {
	name string
	params []string
	public bool
	immutable bool

	block *jade.BlockContext
	function func(args []any, visitor Visitor) any
}

func (f FuncObj) Call(args []any, visitor Visitor) any {
	if f.block != nil {
		return visitor.VisitBlock(f.block)
	} else if f.function != nil {
		return f.function(args, visitor)
	}

	return nil
}

func NewFuncObj(name string, params []string, public bool, immutable bool, block *jade.BlockContext, function func(args []any, visitor Visitor) any) FuncObj {
	return FuncObj{name: name, params: params, public: public, immutable: immutable, block: block, function: function}
}
