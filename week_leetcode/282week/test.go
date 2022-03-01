package main

import "fmt"

func main() {
	// char --> int int()
	// int --> char rune() + string()
	var c rune='a'
	var i int =98
	i1:=int(c) // 好像c 本来存的就是97
	fmt.Println("'a' convert to",i1)
	//c1:=rune(i)
	fmt.Println("98 convert to",string(i))

	for _, char := range []rune("世界你好") {
		fmt.Println(string(char))
	}

	// map operation
	if _, ok := map[key]; ok {
		// 存在
	}

	if _, ok := map[key]; !ok {
		// 不存在
	}
}
