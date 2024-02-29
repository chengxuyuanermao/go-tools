package main

import (
	"fmt"
	"reflect"
)

type c struct {
	name string
}

type d struct {
	name string
}

func main() {
	var a, b, cc, dd interface{}
	a = 1
	b = "1"
	fmt.Println(a == b, reflect.TypeOf(a), reflect.TypeOf(b))

	cc = c{name: "ss"}
	dd = c{name: "dd"}
	fmt.Println(cc == dd)

}
