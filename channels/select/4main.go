package main

import "fmt"

var num = []int{1, 2, 3}

func help(num []int) []int {
	if len(num) == 0 {
		return []int{}
	}

	n := num[len(num)-1]
	num = num[:len(num)-1]
	res := help(num)
	size := len(num)
	for i := 0; i < size; i++ {
		num = append(num[i])
		nu
	}

}

func main() {
	// 谁来的快要谁
	// 因为channel会阻塞 所以如果要非阻塞就可以用select + default
	help(num)
}
