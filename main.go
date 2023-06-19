package main

import (
	"fmt"
	"github.com/chengxuyuanermao/goTools/projectCsv"
	"time"
)

func main() {
	projectCsv.AnalyzeCsvV2()
}

func tickerDemo2() {
	ticker := time.NewTicker(3 * time.Second)
	i := 0
	for {
		now := <-ticker.C

		fmt.Println(now)
		i++
		if i >= 3 {
			ticker.Reset(2 * time.Second)
		}
		if i >= 5 {
			break
		}
	}
	fmt.Println("end")
}

func tickerDemo() {
	tickerChan := time.Tick(2 * time.Second)
	for i := range tickerChan {
		fmt.Println(i)
	}
}

func UseTimer2() {
	timer := time.NewTimer(2 * time.Second)

	// 等待定时器的到期事件
	<-timer.C
	fmt.Println("Timer expired")

	// 重置定时器
	timer.Reset(1 * time.Second)

	// 等待定时器的到期事件
	<-timer.C
	fmt.Println("Timer expired again")

}

func UseTimer() {
	fmt.Println("timer start，", time.Now())
	timer := time.NewTimer(2 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("timer expired")
	}()

	time.Sleep(1 * time.Second)
	if !timer.Stop() {
		select {
		case <-timer.C:
			fmt.Println("stop C ")
		default:
			fmt.Println("default")
		}
	}
	fmt.Println("timer stop")

	isReset := timer.Reset(3 * time.Second)
	if !isReset {
		fmt.Println("false")
		return
	}
	fmt.Println("timer reset")
	nowTime := <-timer.C
	fmt.Println("expired again, ", nowTime)
}

func UseAfter() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	timerChan := time.After(time.Second * 2)
	curTime, xxx := <-timerChan
	fmt.Println(curTime.Format("2006-01-02 15:04:05"), xxx)
}

func ODemo() {
	now := time.Now()
	fmt.Println(now)
	later := now.Add(time.Minute)
	fmt.Println(later)

	res := later.Before(now)
	fmt.Println(res)
}

func addDemo() {
	now := time.Now()
	afterTime := now.Add(time.Minute * 2)
	fmt.Println(afterTime)
}

func subDemo() {
	now := time.Now()
	later := now.Add(time.Minute * 2)
	diff := later.Sub(time.Now())
	fmt.Println(int(diff.Seconds()))
}

func locationDemo() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err)
		return
	}
	timeZone, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-06-19 10:47:00", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("零时区", timeZone)

	locAsia, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	timeAsia := timeZone.In(locAsia)
	fmt.Println(timeAsia)
}
