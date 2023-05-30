package studyReflect

import (
	"fmt"
	"reflect"
)

type cat struct {
	Name string
}

func Use2() {
	var f float64
	f = 3.14
	t1 := reflect.TypeOf(f)
	fmt.Println(t1.String())

	var c cat
	c = cat{
		Name: "cwx",
	}
	t2 := reflect.TypeOf(c)
	fmt.Println(t2.String())
}
