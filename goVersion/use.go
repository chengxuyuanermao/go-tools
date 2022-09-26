package goVersion

import (
	"fmt"
	"github.com/hashicorp/go-version"
)

func Use() {
	v1, err := version.NewVersion("1.2")
	if err != nil {
		fmt.Println(err)
	}
	v2, err := version.NewVersion("1.4+metadata")
	if v1.LessThan(v2) {
		fmt.Println("v1 is less than v2")
	}
}
