package LCR011

// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
func findMaxLength(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	result := 0

	// 先求和，然后再求差
	sum := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		} else {
			nums[i] = 1
		}
		sum += nums[i]
	}

	// 用map存储sum的值，key为sum，value为index
	if sum == 0 {
		result = len(nums)
	}

	left := 0
	right := 0

	for {
		if sum == 0 {
			result = max(result, right-left+1)
		}

		if right == len(nums)-1 {
			break
		}

		right++
		sum -= nums[right]
	}

	return result
}
