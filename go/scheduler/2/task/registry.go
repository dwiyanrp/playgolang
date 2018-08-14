package task

import (
	"fmt"
	"reflect"
	"runtime"
)

type Function interface{}

type FunctionMeta struct {
	Name     string
	function Function
	params   map[string]reflect.Type
}

type FuncRegistry struct {
	funcs map[string]FunctionMeta
}

func NewFuncRegistry() *FuncRegistry {
	return &FuncRegistry{
		funcs: make(map[string]FunctionMeta),
	}
}

func (reg *FuncRegistry) Add(function Function) (FunctionMeta, error) {
	funcValue := reflect.ValueOf(function)
	if funcValue.Kind() != reflect.Func {
		return FunctionMeta{}, fmt.Errorf("Provided function value is not an actual function")
	}

	name := runtime.FuncForPC(funcValue.Pointer()).Name()
	funcInstance, err := reg.Get(name)
	if err == nil {
		return funcInstance, nil
	}
	reg.funcs[name] = FunctionMeta{
		Name:     name,
		function: function,
	}
	return reg.funcs[name], nil
}

func (reg *FuncRegistry) Get(name string) (FunctionMeta, error) {
	function, ok := reg.funcs[name]
	if ok {
		return function, nil
	}
	return FunctionMeta{}, fmt.Errorf("Function %s not found", name)
}
