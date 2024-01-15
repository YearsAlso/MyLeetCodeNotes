package main

import (
	"fmt"
	"github.com/YearsAlso/MyLeetCodeNotes/LCR/LCR008"
	"github.com/YearsAlso/MyLeetCodeNotes/definition"
)

var ListNode definition.ListNode

func main() {
	result := LCR008.Perform(7, []int{2, 3, 1, 2, 4, 3})

	listNode := definition.ListNode{
		Val:  1,
		Next: nil,
	}

	fmt.Print(listNode.Val)
	fmt.Println(result)
}
