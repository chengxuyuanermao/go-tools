package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//docx.Use()
	ca := "ss上我《."
	res := utf8.RuneCountInString(ca)
	fmt.Println(res)
}
