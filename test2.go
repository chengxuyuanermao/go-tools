package main

import (
	"fmt"
	"sort"
)

type a struct {
	s string
}

type byS []*a

func (s byS) Len() int {
	return len(s)
}

func (s byS) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byS) Less(i, j int) bool {
	return s[i].s > s[j].s
}

func test22() {
	d := make([]*a, 0)
	temp := &a{s: "a"}
	temp1 := &a{s: "b"}
	temp2 := &a{s: "c"}
	d = append(d, temp, temp1, temp2)
	fmt.Println("before sort:", d)
	for _, v := range d {
		fmt.Println(v.s)
	}
	sort.Sort(byS(d))
	fmt.Println("after sort:", d)
	for _, v := range d {
		fmt.Println(v.s)
	}

	fmt.Println(d[:1])
}
