package main

import "fmt"

func defangIPaddr(address string) string {

	n, j := len(address), 0
	rets := make([]byte, n+6)
	for i := 0; i < n; i++ {
		if address[i] == '.' {
			rets[j] = '['
			j++
			rets[j] = '.'
			j++
			rets[j] = ']'
			j++
		} else {
			rets[j] = address[i]
			j++
		}
	}
	return string(rets)
}

func main() {
	ret := defangIPaddr("1.1.1.1")
	fmt.Println(ret)

}
