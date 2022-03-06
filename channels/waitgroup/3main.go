package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, c chan int, w worker) {
	for n := range c {
		fmt.Printf("worker % d received %c\n", id, n)
		w.done()
	}

}

type worker struct {
	in chan int
	//done chan bool
	//wg *sync.WaitGroup //再次封装一下 --->有点类似于之前看过的依赖倒置原则？
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		//wg: wg,
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w.in, w)
	// channel也是一等公民
	return w
}
func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker // 给每个worker都创建channel
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20) //20个任务,也可以把wg.Add(1)写在循环里

	for i, worker := range workers {
		worker.in <- 'a' + i
		//wg.Add(1)
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()

}

func main() {
	chanDemo()
}
