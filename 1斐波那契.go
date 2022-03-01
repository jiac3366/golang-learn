package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	p, q := 0, 1
	for i := 2; i <= n; i++ {
		c := p + q
		p = q
		q = c
	}
	return q
}

func main() {
	fmt.Println("hh")
	res := fib(22)
	fmt.Println(res)
}
