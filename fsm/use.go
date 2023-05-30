package fsm

import (
	"fmt"
	"github.com/name5566/leaf/timer"
	"math/rand"
	"time"
)

var c = make(chan bool)

func Test() {
	t := &table{
		fsm:        nil,
		dispatcher: timer.NewDispatcher(1000), // 实例化调度器
		timer:      nil,
	}
	// 初始化新牌局的状态为waiting，且设置好切换每个状态时的调触发函数
	t.fsm = NewFSMLudo(TABLE_STATE_WAITING, t.getFSMEvents())
	// 起一个协程监听延迟函数的执行
	t.start()
	// 触发 wating 状态的执行
	t.fsm.Event(TABLE_STATE_WAITING)

	// chan作为结束信号，防止主进程过早退出
	<-c
}

func (t *table) start() {
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

const (
	TABLE_STATE_NOT = iota
	TABLE_STATE_WAITING
	TABLE_STATE_READY
	TABLE_STATE_PLAY
	TABLE_STATE_RESULT
)

type table struct {
	fsm        *FSM              //当前状态
	dispatcher *timer.Dispatcher // 牌桌计时器
	timer      *timer.Timer      // 牌桌计时器标识
}

func (t *table) waiting() {
	fmt.Println("wating---")
	t.fsm.Event(TABLE_STATE_READY)
	fmt.Println("waiting---ending")
}

func (t *table) ready() {
	fmt.Println("ready---")
	c <- true
	//t.fsm.Event(TABLE_STATE_PLAY)
}

func (t *table) play() {
	fmt.Println("play---")
	t.fsm.Event(TABLE_STATE_RESULT)

}

func (t *table) res() {
	fmt.Println("res---")
	fmt.Println("ending---")
	time.Sleep(time.Second * 2)
	t.fsm.Event(TABLE_STATE_WAITING) // 回到初始点继续循环

	// 50%的概率退出执行
	if rand.Intn(100) < 50 {
		c <- true
	}

}

//状态机
func (t *table) getFSMEvents() []*EventDesc {
	return []*EventDesc{
		{Name: TABLE_STATE_WAITING, Src: []int{TABLE_STATE_WAITING, TABLE_STATE_READY, TABLE_STATE_RESULT}, Enter: func(e *Event) {
			t.runIntervalTime(2, t.waiting)
		}},
		{Name: TABLE_STATE_READY, Src: []int{TABLE_STATE_WAITING, TABLE_STATE_RESULT}, Enter: func(e *Event) {
			t.runIntervalTime(5, t.ready)
		}},
		{Name: TABLE_STATE_PLAY, Src: []int{TABLE_STATE_READY, TABLE_STATE_PLAY}, Enter: func(e *Event) {
			t.play()
		}},
		{Name: TABLE_STATE_RESULT, Src: []int{TABLE_STATE_PLAY}, Enter: func(e *Event) {
			t.res()
		}},
	}
}

//运行计时器
func (t *table) runIntervalTime(t1 int, f func()) {
	t.stopTimer()
	t.timer = t.dispatcher.AfterFunc(time.Duration(t1)*time.Second, func() {
		t.stopTimer()
		f()
	})
}

func (t *table) stopTimer() {
	if t.timer != nil {
		t.timer.Stop()
		t.timer = nil
	}
}
