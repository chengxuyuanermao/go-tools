package sort

import (
	"fmt"
	"sort"
)

func TestInt() {
	a := make([]int, 0)
	a = append(a, 0, 11, 2, 55)

	// 正序排序
	sort.Ints(a)

	fmt.Println(a)
}

//----- 第一种倒序 -----

func TestIntReverse() {
	a := make([]int, 0)
	a = append(a, 0, 11, 2, 55)

	// 降序排序
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j] // i元素更大，且在前面
	})

	fmt.Println(a)
}

//----- 第二种倒序 -----

func TestIntsReverseV2() {
	a := make([]int, 0)
	a = append(a, 0, 11, 2, 55)

	// 降序排序
	sort.Sort(myInts(a))

	fmt.Println(a)
}

type myInts []int

func (m myInts) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m myInts) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m myInts) Len() int {
	return len(m)
}
