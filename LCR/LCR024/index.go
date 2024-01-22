package LCR024

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = definition.ListNode

func reverseList(head *ListNode) *ListNode {
	result := (*ListNode)(nil)

	for head != nil {
		temp := head.Next
		head.Next = result
		result = head
		head = temp
	}

	return result
}
