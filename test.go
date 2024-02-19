package main

import (
	"context"
	"fmt"
	"time"
)

func test2() {
	// 整个方法最多执行5秒，否则超时
	timeOut := 5 * time.Second

	done := make(chan struct{})
	go func() {
		// 模拟业务耗时操作
		time.Sleep(2 * time.Second)
		fmt.Println("goroutine exec finish")
		done <- struct{}{} // 执行完成信号
	}()

	select {
	case <-time.After(timeOut):
		fmt.Println("timeout over") // 超时退出了
	case <-done:
		fmt.Println("done") // 业务完成退出
	}
}

func test3() {
	// 整个方法最多执行5秒，否则超时
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct{})
	go func() {
		// 模拟业务耗时操作
		time.Sleep(2 * time.Second)
		fmt.Println("goroutine exec finish")
		done <- struct{}{} // 执行完成信号
	}()

	select {
	case <-ctx.Done():
		fmt.Println("timeout over") // 超时退出了
	case <-done:
		fmt.Println("done") // 业务完成退出
	}

}
