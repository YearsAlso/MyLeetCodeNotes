package LCR044

import (
	"github.com/YearsAlso/MyLeetCodeNotes/definition"
	"math"
)

type TreeNode = definition.TreeNode

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	result := make([]int, 0)
	stack := []*TreeNode{root}
	nextStack := make([]*TreeNode, 0)

	for {
		maxVal := math.MinInt
		for _, v := range stack {
			if v.Val > maxVal {
				maxVal = v.Val
			}
			if v.Left != nil {
				nextStack = append(nextStack, v.Left)
			}
			if v.Right != nil {
				nextStack = append(nextStack, v.Right)
			}
		}

		result = append(result, maxVal)

		if len(nextStack) == 0 {
			break
		}

		stack = nextStack
		nextStack = make([]*TreeNode, 0)

	}

	return result
}
