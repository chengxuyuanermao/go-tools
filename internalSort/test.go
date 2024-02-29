package internalSort

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
	return s[i].s > s[j].s // 这里为倒序排序;

	/*
		1 看做i元素大于j元素，i元素更大的在数组前面 （即降序）
		2 元素分布顺序: f g h [i j] k l m
	*/
}

func Test22() {
	d := make([]*a, 0)
	temp := &a{s: "a"}
	temp1 := &a{s: "b"}
	temp2 := &a{s: "c"}
	d = append(d, temp, temp1, temp2)
	fmt.Println("before internalSort:")
	for _, v := range d {
		fmt.Println(v.s)
	}

	sort.Sort(byS(d))

	fmt.Println("after internalSort:")
	for _, v := range d {
		fmt.Println(v.s)
	}

}
