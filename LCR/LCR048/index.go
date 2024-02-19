package LCR048

import (
	"fmt"
	"github.com/YearsAlso/MyLeetCodeNotes/definition"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = definition.TreeNode

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	result := make([]string, 0)

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	nextStack := make([]*TreeNode, 0)

	for {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			result = append(result, strconv.Itoa(node.Val))
			nextStack = append(nextStack, node.Right)
			nextStack = append(nextStack, node.Left)
		} else {
			result = append(result, "null")
		}

		stack = nextStack

		nextStack = make([]*TreeNode, 0)

		// 如果栈中全部为nil，则结束
		allNil := false
		for _, v := range stack {
			if v != nil {
				allNil = true
				break
			}
		}
		if !allNil {
			stack = make([]*TreeNode, 0)
		}

		if len(stack) == 0 {
			break
		}
	}

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// 对字符串切分
	dataList := strings.Split(data, ",")

	if len(dataList) == 0 {
		return nil
	}

	if dataList[0] == "null" {
		return nil
	}

	headVal, _ := strconv.Atoi(dataList[0])
	// 递归构建二叉树
	head := &TreeNode{
		Val: headVal,
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, head)

	for i := 1; i < len(dataList); i++ {
		node := stack[len(stack)-1]
		if dataList[i] != "null" {
			val, _ := strconv.Atoi(dataList[i])
			node.Left = &TreeNode{
				Val: val,
			}
			stack = append(stack, node.Left)
		} else {
			node.Left = nil
		}
		i++
		if i < len(dataList) && dataList[i] != "null" {
			val, _ := strconv.Atoi(dataList[i])
			node.Right = &TreeNode{
				Val: val,
			}
			stack = append(stack, node.Right)
		} else {
			node.Right = nil
		}
		stack = stack[:len(stack)-1]
	}

	return head
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func ConstructorTest() {
	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	})
	ans := deser.deserialize(data)

	fmt.Println(data)
	fmt.Println(ans)
}
