package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	arr := [10]int{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				arr[i]++ // 中间协程没有机会切换 而且它不会主动交出控制权
				runtime.Gosched()
			}
		}(i) // 如果不传这个i，内部会引用外部的i,外部如果把i加到10 内部就会发生出错
	}
	time.Sleep(time.Second)
	fmt.Println(arr) // 一边读一边写 会有问题！go run -race .\main.go能检测出来 可以用channel解决
}
