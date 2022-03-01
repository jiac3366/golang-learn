package main

import (
	"fmt"
	"time"
)

func doWorker(id int, c chan int, done chan bool) {
	// 根据1main.go 如何通知外部 我已经把channel接收完了呢？+1个bool型的done channel通知外部
	// 这样的话完整的通信机制需要2个channel
	// 最好把creatWorker的返回值chan int 封装成一个struct
	for n := range c {
		fmt.Printf("worker % d received %d\n", id, n)
		done <- true
	}

}

type worker struct {
	in   chan int
	done chan bool
}

// 返回一个外部只能向其中发数据的channel 所以函数里面只能收数据,chan[]数组的类型也要改成chan<- int
func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	// channel也是一等公民
	return w
}
func chanDemo() {
	var workers [10]worker // 给每个worker都创建channel
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		<-workers[i].done // 如果写在这里 意味着发一个任务 又等着结束
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].done
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
