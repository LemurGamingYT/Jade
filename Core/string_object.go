package Core

type StringObj struct {
	value string
	objType string
}

func (s StringObj) Repr(context any) string {
	return s.value
}

func NewStringObj(value string) StringObj {
	return StringObj{value: value, objType: "string"}
}
