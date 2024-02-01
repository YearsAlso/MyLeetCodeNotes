package definition

// LeetCode TreeNode Definition

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildByArray(array []any) *TreeNode {
	if len(array) == 0 {
		return nil
	}

	stack := make([]*TreeNode, len(array))
	for i, v := range array {
		if v == nil {
			stack[i] = nil
			continue
		}
		elems := &TreeNode{
			Val: v.(int),
		}
		stack[i] = elems
	}

	for i, v := range stack {
		if len(stack) < (i*2 + 2) {
			break
		}

		v.Left = stack[i*2+1]
		v.Right = stack[i*2+2]
	}

	return stack[0]
}

func BuildByIntArray(array []int) *TreeNode {
	if len(array) == 0 {
		return nil
	}

	stack := make([]*TreeNode, len(array))
	for i, v := range array {
		if v == 0 {
			stack[i] = nil
			continue
		}
		elems := &TreeNode{
			Val: v,
		}
		stack[i] = elems
	}

	for i, v := range stack {
		if len(stack) < (i*2 + 2) {
			break
		}

		v.Left = stack[i*2+1]
		v.Right = stack[i*2+2]
	}

	return stack[0]
}

func PreOrderTraverse(root *TreeNode, nodeList []*TreeNode) []*TreeNode {
	stack := []*TreeNode{root}
	nextStack := make([]*TreeNode, 0)
	for {
		for _, v := range stack {
			nodeList = append(nodeList, v)
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
	}

	return nodeList
}
