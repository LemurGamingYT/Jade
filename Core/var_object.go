package Core

type VarObj struct {
	name string
	value any
	constant bool
}

func (v VarObj) Repr(context any) string {
	return "Variable " + v.name
}

func NewVarObj(name string, value any, constant bool) VarObj {
	return VarObj{name: name, value: value, constant: constant}
}
