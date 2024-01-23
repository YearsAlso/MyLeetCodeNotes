package LCR025

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = definition.ListNode

/*
*
给定两个 非空链表 l1和 l2 来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

可以假设除了数字 0 之外，这两个数字都不会以零开头。
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1List := convertToNumber(l1)
	num2List := convertToNumber(l2)

	// 找到最长的那个

	var result *ListNode = nil
	result = nil

	l := max(len(num1List), len(num2List))

	add := 0

	for i := 0; i < l; i++ {
		num := 0
		i1 := len(num1List) - 1 - i
		i2 := len(num2List) - 1 - i
		if i < len(num1List) && i < len(num2List) {
			num = num1List[i1] + num2List[i2] + add
		} else if i < len(num1List) {
			num = num1List[i1] + add
		} else if i < len(num2List) {
			num = num2List[i2] + add
		}
		add = num / 10
		num = num % 10
		result = &ListNode{
			Val:  num,
			Next: result,
		}
	}

	if add != 0 {
		result = &ListNode{
			Val:  add,
			Next: result,
		}
	}

	return result
}

func convertToNumber(l *ListNode) []int {
	result := make([]int, 0)

	for {
		if l == nil {
			break
		}
		result = append(result, l.Val)
		l = l.Next
	}

	return result
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) []int {
	// 转换成数组
	result := addTwoNumbers(l1, l2)

	return definition.GetArray(result)
}
