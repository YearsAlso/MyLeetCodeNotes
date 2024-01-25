package sumIndicesWithKSetBits

/*
给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

请你用整数形式返回 nums 中的特定元素之 和 ，这些特定元素满足：其对应下标的二进制表示中恰存在 k 个置位。

整数的二进制表示中的 1 就是这个整数的 置位 。
*/
func sumIndicesWithKSetBits(nums []int, k int) int {
	result := 0

	flag := 0
	for i := 0; i < k; i++ {
		flag += 1 << i
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < flag {
			continue
		}
		numFlag := flag
		for nums[i] > numFlag {
			if (nums[i] & numFlag) == numFlag {
				result += nums[i]
			}
			numFlag = numFlag << 1
		}
	}

	return result
}
