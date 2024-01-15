package LCR010

// 给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。

// 示例 1 : 使用暴力法
/*func subarraySum(nums []int, k int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]
			if sum == k {
				result++
			}
		}
	}

	return result
}*/

// 示例 2 : 使用滑动窗口
func subarraySum(nums []int, k int) int {
	count := map[int]int{0: 1}
	sum := 0
	result := 0

	for _, num := range nums {
		sum += num
		if _, ok := count[sum-k]; ok {
			result += count[sum-k]
		}
		count[sum]++
	}

	return result
}
