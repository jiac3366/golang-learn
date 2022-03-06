package main

import "fmt"

var res [][]int

func subsets(nums []int) [][]int {
	res = [][]int{}
	backtrack(nums, 0, []int{})
	return res

}

func backtrack(nums []int, level int, path []int) {
	temp := make([]int, len(path))
	copy(temp, path)
	res = append(res, temp)

	for i := level; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrack(nums, i+1, path)
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
