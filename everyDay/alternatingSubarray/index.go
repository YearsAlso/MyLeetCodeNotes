package alternatingSubarray

/*
*
给你一个下标从 0 开始的整数数组 nums 。如果 nums 中长度为 m 的子数组 s 满足以下条件，我们称它是一个 交替子数组 ：

m 大于 1 。
s1 = s0 + 1 。
下标从 0 开始的子数组 s 与数组 [s0, s1, s0, s1,...,s(m-1) % 2] 一样。也就是说，s1 - s0 = 1 ，s2 - s1 = -1 ，s3 - s2 = 1 ，s4 - s3 = -1 ，以此类推，直到 s[m - 1] - s[m - 2] = (-1)m 。
请你返回 nums 中所有 交替 子数组中，最长的长度，如果不存在交替子数组，请你返回 -1 。

子数组是一个数组中一段连续 非空 的元素序列。
*/
func alternatingSubarray(nums []int) int {
	// 对数组做差

	if len(nums) <= 1 {
		return len(nums)
	}

	subList := make([]int, len(nums)-1)

	for i := 1; i < len(nums); i++ {
		subList[i-1] = nums[i] - nums[i-1]
	}

	// 翻转标志位
	flag := false
	result := 0
	maxLen := 0
	for i := 0; i < len(subList); i++ {
		if subList[i] == 1 {
			if flag == false {
				maxLen += 1
				flag = true
			} else {
				result = max(result, maxLen)
				flag = true
				maxLen = 1
			}
		} else if subList[i] == -1 && flag == true {
			maxLen += 1
			flag = false
		} else {
			result = max(result, maxLen)
			maxLen = 0
			flag = false
		}
	}

	if result == 0 && maxLen == 0 {
		return -1
	}

	if maxLen > 0 {
		result = max(result, maxLen) + 1
	} else {
		result += 1
	}

	return result
}
