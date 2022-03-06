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

// 将打印逻辑换成函数
func (node *TreeNode) InorderWithFunc(f func(node *TreeNode)) {
	if node == nil {
		return
	}

	node.Left.InorderWithFunc(f)
	f(node)
	node.Right.InorderWithFunc(f)
}

// 借助channel实现 生成器逻辑
func (node *TreeNode) TravelWithChannel() chan *TreeNode {
	c := make(chan *TreeNode)
	go func() {
		node.InorderWithFunc(func(node *TreeNode) {
			c <- node
		})
		close(c)
	}()
	return c
}
