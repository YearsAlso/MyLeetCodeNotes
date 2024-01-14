package LCR009

import (
	"container/list"
)

// 给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	if len(nums) == 0 {
		return 0
	}

	tmpArray := list.New()
	tmp := 1
	number := 0

	for _, num := range nums {
		tmpArray.PushBack(num)
		tmp *= num

		for tmp >= k {
			tmp /= tmpArray.Remove(tmpArray.Front()).(int)
		}

		number += tmpArray.Len()
	}

	return number
}

/*
* 问：我为什么会犯这种错误，我的编程思想哪方面需要提高
* 答：你的代码逻辑大体上是正确的，你已经理解了滑动窗口的基本概念，并且能够将其应用在实际问题中。然而，你在处理特定情况时，没有完全理解滑动窗口的工作原理，这导致了你的代码不能正确地解决问题。  在你的代码中，你试图找出乘积小于k的连续子数组的个数。你使用了一个滑动窗口（tmpArray）来存储当前的子数组，并使用tmp来存储子数组的乘积。你在每次添加新元素到tmpArray后，都会检查tmp是否大于等于k。如果tmp大于等于k，你就会从tmpArray的头部移除元素，直到tmp小于k。然后，你会增加number（满足条件的子数组的个数）。  这个逻辑在大多数情况下都是正确的，但是在处理tmp大于等于k的情况时，你的代码没有正确地计算子数组的个数。当tmp大于等于k时，你会从tmpArray的头部移除元素，直到tmp小于k。然后，你会增加number。但是，你没有考虑到，即使移除了tmpArray的头部元素，tmpArray中仍然可能存在多个满足条件的子数组。所以，你应该在每次移除tmpArray的头部元素后，都增加number。  这个问题表明，你需要更深入地理解滑动窗口的工作原理，并且需要更仔细地考虑各种可能的情况。在编程时，我们需要考虑所有可能的输入，并确保我们的代码能够正确地处理这些输入。这需要我们对问题有深入的理解，对我们的代码有清晰的理解，以及对我们使用的算法和数据结构有深入的理解。
 */
