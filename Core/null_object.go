package Core

type NullObj struct {
	value string
	objType string
}

func (n NullObj) Repr(context any) string {
	return ReportError("Type", "Cannot read 'null' type")
}

func NewNullObj() NullObj {
	return NullObj{value: "null", objType: "null"}
}
