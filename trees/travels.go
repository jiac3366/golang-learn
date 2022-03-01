package trees

import "fmt"

func (node *TreeNode) Preorder() {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	node.Left.Preorder()
	node.Right.Preorder()
}
