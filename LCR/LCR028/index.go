package LCR028

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

/*
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。
*/

type Node = definition.DoublyLinkedNode

func flatten(root *Node) *Node {

	result := root

	for result != nil {
		if result.Child != nil {
			child := result.Child
			for child.Next != nil {
				child = child.Next
			}
			child.Next = result.Next
			if result.Next != nil {
				result.Next.Prev = child
			}
			result.Next = result.Child
			result.Child.Prev = result
			result.Child = nil
		}
		result = result.Next
	}

	return root

}

func Flatten(root *Node) []int {

	result := flatten(root)

	var list []int

	for result != nil {
		list = append(list, result.Val)
		result = result.Next
	}

	return list
}
