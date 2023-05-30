package studyReflect

import (
	"fmt"
	"reflect"
)

// https://darjun.github.io/2021/05/27/godailylib/reflect/
func Use5() {
	inspectMap(map[uint32]uint32{
		1: 11,
		2: 22,
	})
}

func inspectMap(m interface{}) {
	v := reflect.ValueOf(m)
	for _, k := range v.MapKeys() {
		field := v.MapIndex(k)
		fmt.Printf("%v => %v \n", k.Interface(), field.Interface())
	}
}
