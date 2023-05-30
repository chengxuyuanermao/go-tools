package studyReflect

import (
	"fmt"
	"reflect"
)

/**
透视结构体中定义的方法，需要以下方法：

reflect.Type.NumMethod()：返回结构体定义的方法个数；
reflect.Type.Method(i)：返回第i个方法的reflect.Method对象；

事实上，reflect.Value也定义了NumMethod()/Method(i)这些方法。区别在于：reflect.Type.Method(i)返回的是一个reflect.Method对象，可以获取方法名、类型、是结构体中的第几个方法等信息。如果要通过这个reflect.Method调用方法，必须使用Func字段，而且要传入接收器的reflect.Value作为第一个参数：
	m.Func.Call(v, ...args)
但是reflect.Value.Method(i)返回一个reflect.Value对象，它总是以调用Method(i)方法的reflect.Value作为接收器对象，不需要额外传入。而且直接使用Call()发起方法调用：
	m.Call(...args)
reflect.Type和reflect.Value有不少同名方法，使用时需要注意甄别。
*/

func Use8() {
	u := &User2{}
	inspectMethod(u)
}

func inspectMethod(o interface{}) {
	t := reflect.TypeOf(o)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m)
	}
}

type User2 struct{}

func (this *User2) Test1() {

}
func (this *User2) Test2() {

}
