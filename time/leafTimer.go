package time

import (
	"fmt"
	"github.com/name5566/leaf/timer"
	"time"
)

var c = make(chan bool)

func TestLeafTimer() {
	t := &tableLeaf{
		dispatcher: timer.NewDispatcher(1000), // 实例化调度器
		timer:      nil,
	}

	// 起一个协程监听延迟函数的执行
	t.start()

	// 业务逻辑
	t.runIntervalTime(3, func() {
		t.play()
	})

	// chan作为结束信号，防止主进程过早退出
	<-c
}

func (t *tableLeaf) start() {
	go func() {
		for {
			select {
			case t := <-t.dispatcher.ChanTimer:
				if t != nil {
					t.Cb()
				}
			}
		}
	}()
}

type tableLeaf struct {
	dispatcher *timer.Dispatcher // 牌桌计时器
	timer      *timer.Timer      // 牌桌计时器标识
}

func (t *tableLeaf) play() {
	fmt.Println("play---")
	c <- true
}

//运行计时器
func (t *tableLeaf) runIntervalTime(t1 int, f func()) {
	t.stopTimer()
	t.timer = t.dispatcher.AfterFunc(time.Duration(t1)*time.Second, func() {
		t.stopTimer()
		f()
	})
}

func (t *tableLeaf) stopTimer() {
	if t.timer != nil {
		t.timer.Stop()
		t.timer = nil
	}
}
