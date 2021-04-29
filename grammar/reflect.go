package main

import (
	"fmt"
	"reflect"
)

type T struct {
}

func (T) Add(x int) {
	fmt.Println(x)
}

/**
疑问: 为嘛函数签名必须是大写才可以在同一个包中
	为嘛非要结构体才可以
*/
func main() {
	//var t T
	//params := make([]reflect.Value, 1)
	//params = append(params, reflect.ValueOf(1))
	//reflect.ValueOf(&t).MethodByName("add").Call(params)
	Invoke(T{}, "Add", 1)
}

func Invoke(any interface{}, name string, args ...interface{}) {
	params := make([]reflect.Value, len(args))
	for idx := range args {
		params[idx] = reflect.ValueOf(args[idx])
	}

	reflect.ValueOf(any).MethodByName(name).Call(params)
}
