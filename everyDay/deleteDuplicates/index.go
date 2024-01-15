package deleteDuplicates

import (
	"container/list"
	. "github.com/YearsAlso/MyLeetCodeNotes/definition"
)

func deleteDuplicates(head *ListNode) *ListNode {
	stack := list.New()

	result := &ListNode{
		Val:  head.Val,
		Next: nil,
	}

	for head != nil {
		if head.Val != stack.Back().Value.(int) {
			stack.PushBack(head.Val)
			result.Next = &ListNode{
				Val:  head.Val,
				Next: nil,
			}
		}

		head = head.Next
	}

	return result
}
