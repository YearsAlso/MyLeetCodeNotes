package LCR047

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

type TreeNode = definition.TreeNode

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}

	return root

}
