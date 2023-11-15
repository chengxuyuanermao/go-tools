package pprof

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	if n <= 1 {
		return 1
	}

	f1, f2 := 1, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}

func TestCpu() {
	f, err := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	n := 10
	for i := 1; i <= 5; i++ {
		fmt.Printf("fib(%d)=%d\n", n, fib(n)) // 使用fib2不足10ms就执行完了。运行时每隔 10ms 会中断一次，记录每个 goroutine 当前执行的堆栈，以此来分析耗时
		n += 3 * i
	}
}
