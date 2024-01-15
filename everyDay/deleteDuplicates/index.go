package deleteDuplicates

import (
	"container/list"
	. "github.com/YearsAlso/MyLeetCodeNotes/definition"
)

func deleteDuplicates(head *ListNode) *ListNode {
	stack := list.New()

	result = nil

	for head != nil {
		if stack.Len() == 0 {
			stack.PushBack(head.Val)
			// 创建一个新的节点
			result = &ListNode{
				Val:  head.Val,
				Next: nil,
			}
		} else if head.Val != stack.Back().Value.(int) {
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
