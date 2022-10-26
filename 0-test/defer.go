package test

import (
	"fmt"
	"time"
)

func UseDefer() {
	a := 1
	defer printOut(a) // 此时传进的实参结果为1，输出为1
	go func() {
		// 重新定义了变量a，此为局部变量
		//a := 2
		//_ = a

		// 用的是全局变量a
		a = 2
	}()

	// 延迟了一秒，上面协程会把 a赋值为2，所以下面输出2
	time.Sleep(1 * time.Second)
	defer printOut(a) // 结果为2
}

func UseDefer2() {
	a := 1
	defer func(a int) {
		printOut(a) // 输出1，把a传进闭包函数，确定了a的值
	}(a)
	a = 2
}

func UseDefer3() {
	a := 1
	defer func() {
		printOut(a) // 为2，因为是在闭包函数里，等defer执行才确定a变量
	}()
	// 用的是全局变量a
	a = 2
}

func printOut(a int) {
	fmt.Println(a)
}

/**
而defer即在return执行完，函数退出前执行

defer、return、返回值的执行顺序是：return最先执行，return负责将结果写入返回值中，接着defer开始执行一些收尾工作，最后函数携带当前返回值退出。
*/
func A() (a int) {
	defer func() {
		fmt.Println("defer...")
		a++
	}()
	a = 10
	return
}

func UseDefer4() {
	fmt.Println(A())
}
