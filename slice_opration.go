package main

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := [3]string{"A", "B", "C"}
	b := a[1:3]
	log.Printf("b is %v,len(b) is %d, cap(b) is %d", b, len(b), cap(b))
	// b is [B C],len(b) is 2, cap(b) is 2
	c := b[:]
	log.Printf("c is %v,len(c) is %d, cap(c) is %d", c, len(c), cap(c))
	// c is [B C],len(c) is 2, cap(c) is 2
	d := b[1:]
	log.Printf("d is %v,len(d) is %d, cap(d) is %d", d, len(d), cap(d))
	// d is [C],len(d) is 1, cap(d) is 1
	e := b[:0]
	log.Printf("e is %v,len(e) is %d, cap(e) is %d", e, len(e), cap(e))
	// e is [],len(e) is 0, cap(e) is 2
	f := append(b[:0], b[1:]...) // 相当于将d元素复制到e中(e的容量为2，相当于保留住了原有数组的容量)
	log.Printf("f is %v,len(f) is %d, cap(f) is %d", f, len(f), cap(f))
	g := [6]string{"A", "B", "C", "D", "F", "G"}
	h := append(g[:2], g[3:]...)
	log.Printf("h is %v,len(h) is %d, cap(h) is %d", h, len(h), cap(h))
	//// h is [A B D F G],len(h) is 5, cap(h) is 6
	log.Printf("g is %v,len(g) is %d, cap(g) is %d", g, len(g), cap(g))
	//// 把"C"剔除了 g就变成: g is [A B D F G G],len(g) is 6, cap(g) is 6 为啥???
	//i := append([]string{"C"}, g[2:]...)
	//log.Printf("i is %v,len(i) is %d, cap(i) is %d", i, len(i), cap(i))
	// i is [C C D F G],len(i) is 5, cap(i) is 5
	j := append(g[:2], append([]string{"C"}, g[2:]...)...)
	log.Printf("j is %v,len(j) is %d, cap(j) is %d", j, len(j), cap(j))

}
