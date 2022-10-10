package time

import (
	"fmt"
	"time"
)

// 参考：https://studygolang.com/articles/20640

// 延时两秒执行 time.After(time.Second * 2)
func UseAfter() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	// create a nobuf channel and a goroutine `timer` will write it after 2 seconds
	timeAfterTrigger := time.After(time.Second * 2)

	// will be suspend but we have `timer` so will be not deadlocked
	curTime, _ := <-timeAfterTrigger

	// print current time
	fmt.Println(curTime.Format("2006-01-02 15:04:05"))
}

// 定时执行, 使用分两种场景：执行几次后退出 和 循环执行不退出
// 执行几次就退出的话我们需要回收 time.Ticker
func UseTicker() {
	// 创建一个计时器
	timeTicker := time.NewTicker(time.Second * 2)
	i := 0
	for {
		if i > 5 {
			break
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		i++
		<-timeTicker.C

	}
	// 清理计时器
	timeTicker.Stop()
}

// 循环执行不需要清理的话可以用更简便的time.Tick()方法
func UseTick() {
	// 创建一个计时器
	timeTickerChan := time.Tick(time.Second * 2)
	for {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		<-timeTickerChan
	}
}
