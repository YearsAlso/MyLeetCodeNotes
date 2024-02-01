package LCR043

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

type TreeNode = definition.TreeNode

/*
完全二叉树是每一层（除最后一层外）都是完全填充（即，节点数达到最大，第 n 层有 2n-1 个节点）的，并且所有的节点都尽可能地集中在左侧。

设计一个用完全二叉树初始化的数据结构 CBTInserter，它支持以下几种操作：

CBTInserter(TreeNode root) 使用根节点为 root 的给定树初始化该数据结构；
CBTInserter.insert(int v)  向树中插入一个新节点，节点类型为 TreeNode，值为 v 。使树保持完全二叉树的状态，并返回插入的新节点的父节点的值；
CBTInserter.get_root() 将返回树的根节点。
*/
type CBTInserter struct {
	root     *TreeNode
	nodeList []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	nodeList := preOrderTraverse(root, make([]*TreeNode, 0))

	return CBTInserter{
		root:     root,
		nodeList: nodeList,
	}
}

func preOrderTraverse(root *TreeNode, nodeList []*TreeNode) []*TreeNode {
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

func (this *CBTInserter) Insert(v int) int {
	this.nodeList = append(this.nodeList, &TreeNode{
		Val: v,
	})
	index := len(this.nodeList) - 1

	if index == 0 {
		return this.nodeList[index].Val
	}

	if index%2 == 0 {
		this.nodeList[(index)/2].Right = this.nodeList[index]
		return this.nodeList[(index)/2].Val
	}

	this.nodeList[(index-1)/2].Left = this.nodeList[index]
	return this.nodeList[(index+1)/2].Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

func ConstructorTest(funcName []string, params [][]int) []any {
	var result []interface{}

	var obj CBTInserter
	for i, v := range funcName {
		switch v {
		case "CBTInserter":
			obj = Constructor(definition.BuildByIntArray(params[i]))
			result = append(result, nil)
		case "insert":
			result = append(result, obj.Insert(params[i][0]))
		case "get_root":
			result = append(result, obj.Get_root())
		}
	}

	return result
}
