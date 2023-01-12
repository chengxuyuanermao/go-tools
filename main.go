package main

import (
	"fmt"
	"github.com/chengxuyuanermao/goTools/docx"
	"strings"
)

func main() {
	docx.Use()
	//xx()
}

func xx() {
	filterWord := []rune{'/', '《', '》', ':', '*', '<', '>', '|', '?', '.', '“', '‘'}
	title := "/dd>"
	for _, w := range filterWord {
		fmt.Println("---", string(w))
		fmt.Println(strings.ContainsRune(title, w))
	}

}
