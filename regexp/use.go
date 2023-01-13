package regexp

import (
	"fmt"
	"regexp"
)

func Use() {

	str := "Golang regular expressions example"

	regexp, err := regexp.Compile(`p([a-z]+)e`)
	match := regexp.FindAllStringSubmatch(str, 2)

	fmt.Println("Match: ", match, " Error: ", err)
}
