package LCR027

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

type ListNode = definition.ListNode

/*
给定一个链表的 头节点 head ，请判断其是否为回文链表。

如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。
*/
func isPalindrome(head *ListNode) bool {
	stackList := make([]int, 0)

	for head != nil {
		stackList = append(stackList, head.Val)
		head = head.Next
	}

	start := 0
	end := len(stackList) - 1
	for i := 0; i < len(stackList)/2; i++ {
		if stackList[start] != stackList[end] {
			return false
		}
		start++
		end--
	}

	return true
}
