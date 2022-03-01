package main

func twoSum(nums []int, target int) []int {
	s := map[int]int{}
	for index, num := range nums {
		if val, ok := s[target-num]; ok {
			return []int{val, index}
		}
		s[num] = index
	}

	return nil
}

func twoSum2(nums []int, target int) []int {

	maps := make(map[int]int)
	for i, v := range nums {
		maps[v] = i
	}
	for i, _ := range nums {
		res, ok := maps[target-nums[i]]
		if ok && res != i {
			return []int{i, res}
		}
	}
	return []int{0, 0}
}

func main() {
	lists := []int{2, 7, 11, 15}
	twoSum(lists, 9)

}
