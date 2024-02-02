package LCR045

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

type TreeNode = definition.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	// 使用层序遍历，最后一层的第一个节点就是最左边的节点
	stack := []*TreeNode{root}
	nextStack := make([]*TreeNode, 0)

	for {
		for _, v := range stack {
			if v.Left != nil {
				nextStack = append(nextStack, v.Left)
			}
			if v.Right != nil {
				nextStack = append(nextStack, v.Right)
			}
		}

		if len(nextStack) == 0 {
			break
		}

		stack = nextStack
		nextStack = make([]*TreeNode, 0)
	}

	return stack[0].Val
}
