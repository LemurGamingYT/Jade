package Core

import "math/big"

type IntObj struct {
	value big.Int
	objType string
}

func (i IntObj) Repr(context any) string {
	return i.value.String()
}

func NewIntObj(value big.Int) IntObj {
	return IntObj{value: value, objType: "int"}
}

func NewBigInt(value string) big.Int {
	var v big.Int
	v.SetString(value, 10)
	return v
}
