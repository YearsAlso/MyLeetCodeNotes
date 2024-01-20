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

func EqualList(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}

	if l1 == nil && l2 != nil {
		return false
	}

	if l1 != nil && l2 == nil {
		return false
	}

	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil || l2 != nil {
		return false
	}

	return true
}

func EqualArray(l1 *ListNode, l2 []int) bool {
	if l1 == nil && len(l2) == 0 {
		return true
	}

	if l1 == nil && len(l2) != 0 {
		return false
	}

	if l1 != nil && len(l2) == 0 {
		return false
	}

	current := l1

	for i := 0; i < len(l2); i++ {
		if current.Val != l2[i] {
			return false
		}

		current = current.Next
	}

	if current != nil {
		return false
	}

	return true
}

func GetArray(l *ListNode) []int {
	if l == nil {
		return nil
	}

	result := make([]int, 0)

	current := l

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}
