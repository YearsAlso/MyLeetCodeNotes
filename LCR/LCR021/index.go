package LCR021

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

/*
*
给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/

type ListNode = definition.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	srcHead := head

	// 1. 先遍历一遍，得到链表长度
	l := 0
	for head != nil {
		l++
		head = head.Next
	}

	// 2. 再遍历一遍，删除倒数第 n 个结点
	head = srcHead
	for i := 0; i < l-n-1; i++ {
		head = head.Next
	}

	if l == n {
		srcHead = srcHead.Next
	} else {
		head.Next = head.Next.Next
	}

	return srcHead
}

func RemoveNthFromEnd(head *definition.ListNode, n int) []int {
	result := removeNthFromEnd(head, n)

	return definition.GetArray(result)
}
