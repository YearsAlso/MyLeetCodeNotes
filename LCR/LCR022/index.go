package LCR022

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = definition.ListNode

func detectCycle(head *ListNode) *ListNode {

	nodeMap := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return head
		} else {
			nodeMap[head] = true
		}

		head = head.Next
	}

	return head
}

func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	if slow != fast {
		return nil
	}

	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func DetectCycle(head *definition.ListNode) []int {
	result := detectCycle(head)

	return definition.GetArray(result)
}
