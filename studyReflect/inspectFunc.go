package studyReflect

import (
	"fmt"
	"reflect"
)

/**
透视函数类型，需要以下方法：

reflect.Type.NumIn()：获取函数参数个数；
reflect.Type.In(i)：获取第i个参数的reflect.Type；
reflect.Type.NumOut()：获取函数返回值个数；
reflect.Type.Out(i)：获取第i个返回值的reflect.Type。
*/

func Use7() {
	inspectFunc(add)
}

func add(a, b int) int {
	return a + b
}
func greeting(name string) string {
	return "hello " + name
}

func inspectFunc(m interface{}) {
	t := reflect.TypeOf(m)
	fmt.Println("input:")
	for i := 0; i < t.NumIn(); i++ {
		fmt.Printf("%v ", t.In(i))
	}

	fmt.Println("\noutput:")
	for i := 0; i < t.NumOut(); i++ {
		fmt.Printf("%v ", t.Out(i))
	}
}
