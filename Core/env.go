package Core

import (
	"math"
	"strconv"
)

type Env struct {
	variables map[string]VarObj
	functions map[string]FuncObj
}

func NewEnv() Env {
	variables := make(map[string]VarObj)
	variables["Pi"] = NewVarObj("Pi", NewFloatObj(NewBigFloat(strconv.FormatFloat(math.Pi, 'f', 4, 64))), true)

	functions := make(map[string]FuncObj)
	functions["Println"] = NewFuncObj("Println", []string{"Value"}, false, true, nil, Println)

	return Env{variables: variables, functions: functions}
}
