package trees

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node TreeNode) Print() {
	fmt.Print(node.Val, " ")
}

func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Val = value
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{Val: value}
}
