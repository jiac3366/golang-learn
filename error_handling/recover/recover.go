package main

import (
	"fmt"
)

func tryRecover() {
	// recover()仅在defer中调用 并且可以获取panic的值 如果无法处理，比如panic(123)（程序都不知道是什么）可以重新panic
	defer func() {
		// 可以获取panic的值
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error be recover", err)
		} else {
			// 再次panic
			//panic(r)
			panic(fmt.Sprintf(
				"I don't know what to do: %v", r))
		}
	}()
	//panic(errors.New("this is new error"))
	panic(111)
	//panic: 111 [recovered]
	//panic: 111

}

func main() {
	tryRecover()

}
