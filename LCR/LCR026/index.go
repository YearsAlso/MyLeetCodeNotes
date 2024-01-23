package LCR026

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

type ListNode = definition.ListNode

func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}

	// 只有两位也不需要变化
	if head.Next.Next == nil {
		return
	}

	listMap := make(map[int]*ListNode)

	l := 0

	for head != nil {
		listMap[l] = head
		l += 1
		head = head.Next
	}

	var result *ListNode = nil
	var last *ListNode = nil

	// 这里循环的条件是 i < l，而不是 i < l/2，因为如果是 i < l/2 的话，那么在 l 为奇数的时候，会出现最后一个节点没有被处理的情况
	for i := 0; i < l/2; i++ {
		if result == nil {
			result = listMap[i]
			last = listMap[i]
		} else {
			last.Next = listMap[i]
			last = last.Next
		}

		if i != l-1-i {
			last.Next = listMap[l-1-i]
			last = last.Next
		}
	}

	if l%2 == 1 {
		last.Next = listMap[l/2]
		last = last.Next
	}

	if last != nil {
		last.Next = nil
	}

	head = result
}

func ReorderList(head *ListNode) []int {
	reorderList(head)
	return definition.GetArray(head)
}
