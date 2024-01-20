package LCR023

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = definition.ListNode

// 给定两个单链表的头节点 headA 和 headB ，请找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)

	for headA != nil {
		nodeMap[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := nodeMap[headB]; ok {
			return headB
		}

		headB = headB.Next
	}

	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	// 使用双指针
	for headA != headB {
		if headA == nil {
			headA = headB
		} else {
			headA = headA.Next
		}
		if headB == nil {
			headB = headA
		} else {
			headB = headB.Next
		}
	}

	return headA

}
