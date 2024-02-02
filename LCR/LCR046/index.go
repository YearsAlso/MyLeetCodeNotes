package LCR046

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

type TreeNode = definition.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if i == size-1 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}

	return res
}
