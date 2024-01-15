package LCR011

// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
/*func myFindMaxLength(nums []int) int {
	if len(nums) <= 0 {
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
		return result
	}

	left := 0
	right := len(nums) - 1

	for {
		if sum == 0 {
			result = max(result, right-left+1)
			break
		}

		if left >= right {
			break
		}

		if sum > 0 {
			if nums[left] == 1 {
				sum -= nums[left]
				left++
			}
			if nums[right] == 1 {
				sum -= nums[right]
				right--
			}

			if nums[left] == -1 && nums[right] == -1 {
				sum += 2
				left++
				right--
			}
		} else {
			if nums[left] == -1 {
				sum -= nums[left]
				left++
			}
			if nums[right] == -1 {
				sum -= nums[right]
				right--
			}

			// 还有一种情况，就是sum小于0，但是left和right都是1，这时候就需要把left和right都向中间移动
			if nums[left] == 1 && nums[right] == 1 {
				sum -= 2
				left++
				right--
			}
		}

	}

	return result
}*/

/*
*
这段代码的逻辑主要是在处理二进制数组，寻找含有相同数量的 0 和 1 的最长连续子数组。代码中使用了两个指针，left 和 right，从数组的两端向中间移动，同时根据 sum 的值来调整 left 和 right 的位置。

然而，这段代码可能存在一些问题：

1. 在处理 sum > 0 和 sum < 0 的情况时，代码中的逻辑可能会导致 left 和 right 指针的移动过快，从而跳过了可能的最长连续子数组。例如，当 nums[left] == 1 和 nums[right] == 1 时，无论 sum 的值是大于0还是小于0，都会使 left 和 right 同时向中间移动，这可能会导致跳过某些有效的子数组。

2. 代码中没有处理 sum == 0 的情况。当 sum == 0 时，理论上应该已经找到了一个含有相同数量的 0 和 1 的子数组，但是代码中并没有对这种情况进行处理。

3. 代码中的循环条件可能会导致无限循环。如果数组中的 0 和 1 的数量不相等，那么 sum 的值可能永远无法达到0，从而导致无限循环。

为了解决这些问题，你可以尝试使用哈希表来存储每个 sum 值第一次出现的位置，然后在遍历数组的过程中，如果当前的 sum 值已经在哈希表中存在，那么就可以计算出一个可能的子数组长度，然后更新结果。这种方法的时间复杂度为 O(n)，空间复杂度也为 O(n)。

这段代码首先将数组中的 0 替换为 -1，然后使用一个哈希表 sumIndexMap 来存储每个 sum 值第一次出现的位置。在遍历数组的过程中，如果当前的 sum 值已经在哈希表中存在，那么就可以计算出一个可能的子数组长度，然后更新结果。如果当前的 sum 值在哈希表中不存在，那么就将其添加到哈希表中。
*/
func findMaxLength(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	maxLength, sum := 0, 0
	sumIndexMap := map[int]int{0: -1}

	for i := 0; i < n; i++ {
		sum += nums[i]
		if prevIndex, has := sumIndexMap[sum]; has {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			sumIndexMap[sum] = i
		}
	}

	return maxLength
}
