package definition

// LeetCode ListNode Definition

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}

	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{
			Val:  nums[i],
			Next: nil,
		}

		current = current.Next
	}

	return head
}
