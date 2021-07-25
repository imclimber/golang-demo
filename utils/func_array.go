package utils

import (
	"errors"
	"reflect"
)

// FuncArray ...
type FuncArray []reflect.Value

// Add ...
func (f *FuncArray) Add(fn interface{}) {

	v := reflect.ValueOf(fn)
	*f = append(*f, v)

}

// Call ...
func (f *FuncArray) Call(index int, params ...interface{}) (result []reflect.Value, err error) {
	if index < 0 || index >= len(*f) {
		err = errors.New("out of index")
		return
	}

	// 控制入参数准确
	if len(params) != (*f)[index].Type().NumIn() {
		err = errors.New("error params")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = (*f)[index].Call(in)
	return
}
