package studyReflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Name    string
	Age     int
	Married bool
}

func Use3() {
	u := &User{
		Name:    "ccc",
		Age:     18,
		Married: false,
	}
	inspectStruct(u)
}

func inspectStruct(m interface{}) {
	v := reflect.ValueOf(m).Elem() // elem进行解指针
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Printf("field:%d type:%s value:%d\n", i, field.Type().Name(), field.Int()) // field.Type().Name() 可获取到自定义的类型 main.User 而不是struct

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fmt.Printf("field:%d type:%s value:%d\n", i, field.Type().Name(), field.Uint())

		case reflect.Bool:
			fmt.Printf("field:%d type:%s value:%t\n", i, field.Type().Name(), field.Bool())

		case reflect.String:
			fmt.Printf("field:%d type:%s value:%q\n", i, field.Type().Name(), field.String())

		default:
			fmt.Printf("field:%d unhandled kind:%s\n", i, field.Kind())
		}
	}
}
