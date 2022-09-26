package goSpew

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Project struct {
	Name string
}

func Use() {
	f := []*Project{
		&Project{Name: "张三"},
		&Project{Name: "李四"},
	}

	// 可以打印出指针里的内容
	spew.Config.Indent = "\t"
	spew.Dump(f)
	// 只能打印出指针
	fmt.Printf("f:%+v\n", f)
}
