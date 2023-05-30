package studyReflect

import (
	"fmt"
	"reflect"
)

func Use9() {
	invoke(add, 1, 2)
	invoke(greeting, "dj")
}

func invoke(f interface{}, args ...interface{}) {
	fmt.Println("------")
	v := reflect.ValueOf(f)

	argV := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		argV = append(argV, reflect.ValueOf(arg))
	}

	fmt.Println("outPut:")
	rets := v.Call(argV)
	for _, ret := range rets {
		fmt.Println(ret.Interface())
	}
}
