package Core

import "strconv"

type BoolObj struct {
	value bool
	objType string
}

func (b BoolObj) Repr(context any) string {
	return strconv.FormatBool(b.value)
}

func NewBoolObj(value bool) BoolObj {
	return BoolObj{value: value, objType: "bool"}
}

func NewBool(value string) bool {
	var v bool
	v, _ = strconv.ParseBool(value)
	return v
}
