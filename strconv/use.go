package strconv

import (
	"fmt"
	"strconv"
)

func UseStrconv() {
	// 字符串转成int
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}

	// int转成字符串
	a := 11
	b := strconv.Itoa(a)
	fmt.Printf("type:%T, value:%v \n", b, b)
	c := string(65)                          // 变成ascii码
	fmt.Printf("type:%T, value:%v \n", c, c) // type:string, value:A

	// Parse类函数
	// Parse类函数用于转换 字符串为给定类型的值. 类型转换：字符串->all
	bv, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-2", 10, 64)
	u, _ := strconv.ParseUint("2", 10, 64)
	fmt.Println(bv, f, i, u)

	// Format系列函数
	// Format系列函数用于转换 字符串为给定类型的值.类型转换： all->string
	s11 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)
	s4 := strconv.FormatUint(2, 16)
	fmt.Println(s11, s2, s3, s4)
}
