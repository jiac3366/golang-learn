package main

import (
	"fmt"
	"strings"
)

func prefixCount(words []string, pref string) int {
	count := 0
	for _, item := range words {
		res := strings.HasPrefix(item, pref)
		if res {
			count++
		}
	}
	return count
}
func main() {

	//words := []string{"pay","attention","practice","attend"}
	words := []string{"leetcode", "win", "loops", "success"}
	//pref := "at"
	pref := "code"
	fmt.Println(prefixCount(words, pref))

}
