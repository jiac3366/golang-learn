package main

import (
	"fmt"
	"programming/trees"
)

func main() {
	node := trees.CreateNode(3)
	node.Left = &trees.TreeNode{}
	node.Right = &trees.TreeNode{Val: 4}
	node.Right.Left = &trees.TreeNode{Val: 8}
	node.Right.Right = &trees.TreeNode{Val: 9}
	node.Right.Right.Left = &trees.TreeNode{Val: 119}
	// 体现函数式编程--求树节点个数
	treeAmount := 0
	node.InorderWithFunc(func(n *trees.TreeNode) {
		fmt.Println(n.Val)
		treeAmount++
	})
	fmt.Println("个数：", treeAmount)

	maxNode := 0
	c := node.TravelWithChannel()
	for node := range c {
		if node.Val > maxNode {
			maxNode = node.Val
		}
	}

	fmt.Println("MaxNode", maxNode)
}
