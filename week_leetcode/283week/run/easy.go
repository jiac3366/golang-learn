package main

import "fmt"

func cellsInRange(s string) []string {
	var res []string
	for i := s[0]; i <= s[3]; i++ {
		for j := s[1]; j <= s[4]; j++ {
			res = append(res, string(i)+string(j))
		}
	}
	return res
}

func main() {
	s := "K1:Z2"
	fmt.Println(cellsInRange(s))

}
