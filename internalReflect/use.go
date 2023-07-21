package internalReflect

import (
	"fmt"
	"reflect"
	"strings"
)

// Test 参考：https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247490308&idx=1&sn=368533ff3d450c8003a7e81d42cd0bc5&scene=21#wechat_redirect
func Test() {
	a := A{
		Name: "hhhh",
		Age:  10,
	}
	test(a)
}

func test(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Println(t, v, t.Kind())

	// 获取结构体中的字段内容和值
	fmt.Println("获取结构体中的字段内容和值-----")
	table := t.Name()
	sql := fmt.Sprintf("%s%s", "insert into ", table)
	fieldStr := fmt.Sprintf("%s", "(")
	fieldVal := fmt.Sprintf("%s", " values (")
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, t.Field(i).Type, v.Field(i))

		fieldStr += t.Field(i).Name + ","

		switch t.Field(i).Type.Kind() {
		case reflect.String:
			fieldVal += "`" + v.Field(i).String() + "`,"
		case reflect.Int:
			fieldVal += fmt.Sprintf("%d,", v.Field(i).Int())
		default:
			panic("not match error")
		}
	}

	fieldStr = strings.TrimRight(fieldStr, ",") + ")"
	fieldVal = strings.TrimRight(fieldVal, ",") + ")"
	sql += fieldStr + fieldVal
	fmt.Println(sql)

}

type A struct {
	Name string
	Age  int
}

type B struct {
	Id  int
	Sex string
}
