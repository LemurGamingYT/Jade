package Core

import "math/big"

type FloatObj struct {
	value big.Float
	objType string
}

func (f FloatObj) Repr(context any) string {
	return f.value.String()
}

func NewFloatObj(value big.Float) FloatObj {
	return FloatObj{value: value, objType: "float"}
}

func NewBigFloat(value string) big.Float {
	var v big.Float
	v.SetString(value)
	return v
}
