package main

import (
	"fmt"
	"strings"
)

func main() {
	//docx.Use()
	//redBookSpider.Use()
	xx()
}

func xx() {
	res := strings.Contains("中国  广东  广州", "广州")
	fmt.Println(res)

}
