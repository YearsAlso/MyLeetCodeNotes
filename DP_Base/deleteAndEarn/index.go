package deleteAndEarn

import "sort"

/*
给你一个整数数组 nums ，你可以对它进行一些操作。

每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。

开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
*/
func deleteAndEarn(nums []int) int {
	numMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = numMap[nums[i]] + nums[i]
	}

	// 将map 的key 转换为数组，并且排序
	nums = []int{}
	for k := range numMap {
		nums = append(nums, k)
	}

	// 排序
	sort.Ints(nums)

	prev1, prev2 := 0, numMap[nums[0]]

	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		temp := prev2
		prev2 = max(prev1+numMap[i], prev2)
		prev1 = temp
	}

	return max(prev1, prev2)
}
