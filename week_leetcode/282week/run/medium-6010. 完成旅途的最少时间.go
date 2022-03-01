package main

import "fmt"

func minimumTime(time []int, totalTrips int) int64 {
	curTrips := 0
	flag := true
	costime := 0
	timeMap := map[int]int{}
	curtime := 0
	for _, item := range time {
		timeMap[item] = item
	}
	for flag {
		if curTrips >= totalTrips {
			break
		}
		curtime++
		for i := 1; i <= curtime; i++ {
			if _, ok := timeMap[i]; ok {
				curTrips += timeMap[i]
				costime += timeMap[i]
			}
		}
	}
	return int64(costime)
}

func main() {

	fmt.Println(minimumTime([]int{5, 10, 10}, 9))
}
