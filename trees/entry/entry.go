package main

import "programming/trees"

func main() {
	node := trees.CreateNode(3)
	node.Left = &trees.TreeNode{}
	node.Right = &trees.TreeNode{Val: 4}
	node.Right.Left = &trees.TreeNode{Val: 8}
	node.Right.Right = &trees.TreeNode{Val: 9}
	node.Preorder()
}
