package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	//收到close()后停止接收的2种方式
	//1、
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker % d received %d\n", id, n)
	//}

	//2、不用显式close()也可以接收，直到main函数才停止，如何提前停止？2main.go优化这一点
	for n := range c {
		fmt.Printf("worker % d received %d\n", id, n)
	}

}

// 自建 channel
// 返回一个外部只能向其中发数据的channel 所以函数里面只能收数据,chan[]数组的类型也要改成chan<- int
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	// channel也是一等公民
	return c
}
func chanDemo() {
	//var c chan int  // select的时候会用到
	var channels [10]chan<- int // 给每个worker都创建channel
	for i := 0; i < 10; i++ {
		//1、
		//channels[i] = make(chan int)
		//go func(id int, c chan int) {
		//	for {
		//		fmt.Printf("worker % d received %c\n", id, <-c)  // 把%d-->变 成%c
		//	}
		//}(i, channels[i])  // channels[i]
		//2、
		channels[i] = createWorker(i)
	}

	//c <- 1 //发了数据要有人接收 channel用于goroutine之间的交互
	//c <- 2

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	// 一定是发送方close
	close(c) // 一旦close()之后 收到的永远是“零值” 收1ms
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	bufferChannel()
}
