package main

import (
	"container/heap"
	"fmt"
)

//nums = [1,1,1,2,2,3], k = 2
func main() {
	//input:

	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	s := map[int]int{}
	for _, num := range nums {
		// skill:map如何判断key是否存在  s[num] == false
		s[num]++
	}
	//

	h := &IntHeap{}
	heap.Init(h)
	for key, v := range s {
		heap.Push(h, [2]int{key, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	fmt.Print(ret)
	//return ret

}

type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//// This example inserts several ints into an IntHeap, checks the minimum,
//// and removes them in order of priority.
//func main() {
//	h := &heaps.IntHeap{100,16,4,8,70,2,36,22,5,12}
//
//	fmt.Println("\nHeap:")
//	heap.Init(h)
//
//	fmt.Printf("最小值: %d\n", (*h)[0])
//
//	//for(Pop)依次输出最小值,则相当于执行了HeapSort
//	fmt.Println("\nHeap sort:")
//	for h.Len() > 0 {
//		fmt.Printf("%d ", heap.Pop(h))
//	}
//
//	//增加一个新值,然后输出看看
//	fmt.Println("\nPush(h, 3),然后输出堆看看:")
//	heap.Push(h, 3)
//	for h.Len() > 0 {
//		fmt.Printf("%d ", heap.Pop(h))
//	}
//
//
//	fmt.Println("\n使用sort.Sort排序:")
//	h2 := heaps.IntHeap{100,16,4,8,70,2,36,22,5,12}
//	sort.Sort(h2)
//	for _,v := range h2 {
//		fmt.Printf("%d ",v)
//	}
//}

/*
输出结果:
-----------------------------------
Heap:
最小值: 2

Heap sort:
2 4 5 8 12 16 22 36 70 100
Push(h, 3),然后输出堆看看:
3
使用sort.Sort排序:
2 4 5 8 12 16 22 36 70 100
*/
