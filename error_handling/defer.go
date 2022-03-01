package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	//defer的作用 不怕中间出现return 或者 panic 而且defer是按照栈逆序输出
	defer fmt.Println(1)
	fmt.Println(2)
	defer fmt.Println(3)
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	// 自定义error
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err) // 如果不是PathError, 真不知道怎么办了
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err) //open 111.txt The file exists.
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintln(writer, "ss")

}
func main() {
	//tryDefer()
	//f := fib.Fibonacci()
	writeFile("111.txt")
}
