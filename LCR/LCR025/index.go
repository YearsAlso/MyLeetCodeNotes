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

	for i = l - 1; i >= 0; i-- {
		num := 0
		if i < len(num1List) && i < len(num2List) {
			num = num1List[i] + num2List[i] + add
			add = num / 10
		} else if i < len(num1List) {
			num = num1List[i] + add
			add = num / 10
		} else if i < len(num2List) {
			num = num2List[i] + add
			add = num / 10
		}
		result = &ListNode{
			Val:  num,
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
