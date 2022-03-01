package main

import "fmt"

//var arr []int
func minSteps(s string, t string) int {
	arr := make([]int, 26)
	sameCount := 0
	for _, ch := range s {
		arr[int(ch)-97] += 1
	}
	for _, ch := range t {
		if arr[int(ch)-97] != 0 {
			arr[int(ch)-97] -= 1
			sameCount++
		}
	}
	fmt.Println(sameCount)
	return (len(s) - sameCount) + (len(t) - sameCount)
}

func main() {
	s := "leettcode"
	t := "cotas"
	//s := "cotxazilut"
	//t :="nahrrmcchxwrieqqdwdpneitkxgnt"

	fmt.Println(minSteps(s, t))

}
