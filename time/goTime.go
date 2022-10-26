package time

import (
	"fmt"
	"time"
)

/**
https://studygolang.com/articles/26473
*/

func UseGoTime() {
	//now := time.Now()
	//fmt.Printf("current time : %v \n", now)
	//
	//year := now.Year()
	//month := now.Month()
	//day := now.Day()
	//hour := now.Hour()
	//minute := now.Minute()
	//second := now.Second()
	//fmt.Println(year, month, day, hour, minute, second)

	locationDemo()
}

// timezoneDemo 时区示例
func addDemo() {
	now := time.Now()
	fmt.Println(now)
	later := now.Add(time.Minute)
	fmt.Println(later)
}

func subDemo() {
	now := time.Now()
	fmt.Println(now)
	later := now.Add(time.Minute)
	fmt.Println(later)
	diff := later.Sub(now)
	fmt.Println(diff)
}

func ODemo() {
	now := time.Now()
	fmt.Println(now)
	later := now.Add(time.Minute)
	fmt.Println(later)

	res := later.After(now)
	fmt.Println(res)
}

func tickDemo() {
	ticker := time.Tick(2 * time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}

func locationDemo() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(now.Sub(timeObj))
}
