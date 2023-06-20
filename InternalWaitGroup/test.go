package InternalWaitGroup

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup

func Test() {
	wg = &sync.WaitGroup{}
	wg.Add(3)

	go worker(1)
	go worker(2)
	go worker(3)

	wg.Wait()
	fmt.Println("main exit")
}

func worker(i int) {
	defer wg.Done()
	fmt.Println("worker", i)
}
