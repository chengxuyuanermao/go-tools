package time

import (
	"fmt"
	"time"
)

func BlockMain() {
	var nums int = 0
	var c = make(chan bool)
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("begin------------")
			time.Sleep(10 * time.Second)
			nums++         // 外部变量
			if nums == 2 { // 累加到2之后（协程开到3个），发送信号让主进程退出
				c <- true
			}
		}()
	}
	// 阻塞等待退出
	<-c
}

func Test() {
	t := &table{
		TList: make([]*Timer, 0),
	}
	c := make(chan bool)

	// 起协程进行监听
	t.run()

	// 延迟执行
	fmt.Println(time.Now())
	t.AddUniueTimer(999, 5, t.testTimer, &c)

	// 防止主进程过早退出
	//time.Sleep(time.Second * 10)
	<-c
}

func (t *table) testTimer(d interface{}) {
	fmt.Println("11----")
	fmt.Println(time.Now())

	sign, ok := d.(*chan bool)
	if ok {
		*sign <- true
	} else {
		fmt.Println("assert error--")
	}
	// 或者直接传chan（本身为指针类型）
	//sign, ok := d.(chan bool)
	//if ok {
	//	sign <- true
	//} else {
	//	fmt.Println("assert error--")
	//}
}

func (t *table) run() {
	go func() {
		tNew := time.NewTicker(time.Second)
		for {
			select {
			case <-tNew.C:
				t.DoTimer()
			}
		}
	}()
}

type table struct {
	TList []*Timer // 秒定时器列表
}

type Timer struct {
	Id int               // id
	T  int               //定时时间
	H  func(interface{}) // 执行的方法
	D  interface{}       // 方法的参数
}

//定时器
func (this *table) DoTimer() {
	// @end
	if len(this.TList) == 0 {
		return
	}
	nlist := []*Timer{}
	olist := []*Timer{}
	for _, v := range this.TList {
		v.T--         // 由于前面是定时每秒执行，这里减去一秒
		if v.T <= 0 { // 说明到时间执行了
			olist = append(olist, v)
		} else {
			nlist = append(nlist, v)
		}
	}
	this.TList = nlist
	for _, v := range olist {
		v.H(v.D) // 执行方法
	}
}

// AddTimer id,时间,执行的方法,方法的参数
func (this *table) AddTimer(id int, t int, h func(interface{}), d interface{}) {
	this.TList = append(this.TList, &Timer{
		Id: id,
		H:  h,
		T:  t,
		D:  d,
	})
}

// AddUniueTimer 同一id的定时器只能存在一个
// id,时间,执行的方法,方法的参数
func (this *table) AddUniueTimer(id int, t int, h func(interface{}), d interface{}) {
	for i := len(this.TList) - 1; i >= 0; i-- {
		if this.TList[i].Id == id {
			// 此前存在同一id的timer器，先删除
			this.TList = append(this.TList[:i], this.TList[i+1:]...)
		}
	}
	// 把新的timer加上去
	this.AddTimer(id, t, h, d)
}

func (this *table) DelTimer(id int) {
	for i, v := range this.TList {
		if v.Id == id {
			this.TList = append(this.TList[:i], this.TList[i+1:]...)
			break
		}
	}
}

func (this *table) GetTimerNum(id int) int {
	for _, v := range this.TList {
		if v.Id == id {
			return v.T
		}
	}
	return 0
}

//清空定时器
func (this *table) ClearTimer() {
	this.TList = []*Timer{}
}
