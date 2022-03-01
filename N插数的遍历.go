package main

type Node struct {
	Val      int
	Children []*Node
}

var res []int

func preorder(root *Node) []int {
	res := []int{}
	helper(root)
	return res
}

func helper(root *Node) {
	if root == nil {
		res := append(res, root.Val)
		for _, node := range root.Children {
			helper(node)
		}
	}
}
