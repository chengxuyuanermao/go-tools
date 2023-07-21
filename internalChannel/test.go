package internalChannel

import (
	"fmt"
	"sync"
	"time"
)

//参考： https://blog.csdn.net/tool007/article/details/124329558

func Test() {
	//define()
	test5()
}

// -----------
func test5() {
	ch := make(chan int)
	quit := make(chan struct{})
	go func() {
		for i := 0; i < 4; i++ { // 控制fib的次数
			<-ch
		}
		quit <- struct{}{}
	}()
	fib(ch, quit)
}

func fib(ch chan int, quit chan struct{}) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("ending")
			fmt.Println("res : ", y)
			return
		}
	}
}

// -----------
func test4() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("for ending, ", ok)
			break
		}
		fmt.Println("get val ", val)
	}
	fmt.Println("ending")
}

// -----------
func test3() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)

}

// -----------
func define() {
	// 只读channel
	//var readOnlyChannel <-chan int

	// 只写channel
	//var writeOnlyChannel chan<- int

	// 可读可写channel
	var ch chan int
	ch = make(chan int)
	go func() {
		ch <- 3
	}()
	time.Sleep(time.Second)
	fmt.Println(<-ch)
}

// -----------
func test1() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	wg.Add(1)
	go send(ch)
	go rece(ch, wg)

	wg.Wait()
	fmt.Println("test1")
}

func send(ch chan<- int) { // 定义成只写通道
	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Println("ch closing")
	close(ch) // 及时close，可通知rece中遍历的channel已经结束，防止程序死锁
}

func rece(ch <-chan int, wg *sync.WaitGroup) { // 定义成只读通道
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("ch close")
	wg.Done()
}

// -----------

func test2() {
	ch := make(chan int)
	go sendV2(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("ending")
}
func sendV2(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

// -----------
