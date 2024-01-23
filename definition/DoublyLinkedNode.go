package definition

type DoublyLinkedNode struct {
	Val   int
	Prev  *DoublyLinkedNode
	Next  *DoublyLinkedNode
	Child *DoublyLinkedNode
}

func BuildDoublyLinkedNode(nums []int) *DoublyLinkedNode {
	if len(nums) == 0 {
		return nil
	}

	head := &DoublyLinkedNode{
		Val:  nums[0],
		Next: nil,
	}

	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &DoublyLinkedNode{
			Val:  nums[i],
			Next: nil,
		}

		current = current.Next
	}

	return head
}
