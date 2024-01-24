package LCR029

import "github.com/YearsAlso/MyLeetCodeNotes/definition"

type Node = definition.LoopLinkedNode

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		aNode = &Node{
			Val:  x,
			Next: nil,
		}
		aNode.Next = aNode
		return aNode
	}

	current := aNode
	for current.Next != aNode {
		if current.Val <= x && current.Next.Val >= x {
			break
		}

		if current.Val > current.Next.Val && (current.Val <= x || current.Next.Val >= x) {
			break
		}

		current = current.Next
	}

	current.Next = &Node{
		Val:  x,
		Next: current.Next,
	}

	return aNode
}
