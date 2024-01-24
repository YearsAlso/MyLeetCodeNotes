package maximumSumOfHeights

import "fmt"

/*
给你一个长度为 n 下标从 0 开始的整数数组 maxHeights 。

你的任务是在坐标轴上建 n 座塔。第 i 座塔的下标为 i ，高度为 heights[i] 。

如果以下条件满足，我们称这些塔是 美丽 的：

1 <= heights[i] <= maxHeights[i]
heights 是一个 山脉 数组。
如果存在下标 i 满足以下条件，那么我们称数组 heights 是一个 山脉 数组：

对于所有 0 < j <= i ，都有 heights[j - 1] <= heights[j]
对于所有 i <= k < n - 1 ，都有 heights[k + 1] <= heights[k]
请你返回满足 美丽塔 要求的方案中，高度和的最大值 。
*/

// 时间复杂度：O(n^2)，可以正常通过，但是时间复杂度太高，会超时
func maximumSumOfHeights(maxHeights []int) int64 {
	result := int64(0)
	l := len(maxHeights)
	results := make(chan int64, l)

	for i := 0; i < l; i++ {
		if int64(l*maxHeights[i]) <= result {
			continue
		}
		go func(i int) {
			res := getResultByIndex(maxHeights, i)
			results <- res
		}(i)
	}

	for i := 0; i < l; i++ {
		res := <-results
		if res > result {
			result = res
		}
	}

	return result
}

func getResultByIndex(srcHeights []int, maxIndex int) int64 {
	maxValue := srcHeights[maxIndex]
	result := int64(maxValue)

	pre := maxValue
	for i := maxIndex - 1; i >= 0; i-- {
		pre = min(pre, srcHeights[i])
		fmt.Print(pre)
		result += int64(pre)
	}

	suf := maxValue
	for i := maxIndex + 1; i < len(srcHeights); i++ {
		suf = min(srcHeights[i], suf)
		fmt.Print(suf)
		result += int64(suf)
	}

	fmt.Println(result)
	return result
}
