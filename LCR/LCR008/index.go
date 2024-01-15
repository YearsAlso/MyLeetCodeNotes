package LCR008

import (
	"container/list"
	"fmt"
	"github.com/YearsAlso/MyLeetCodeNotes/utils"
	"strconv"
)

// 没有用的函数，这个函数当初的想法是，获得最小的子数组，然后将这个子数组从原数组中删除，然后继续寻找最小的子数组，但是没有动态的比较长度，所以导致错误
func refreshArray(
	tmpArray *list.List,
	tmpSum *int,
	target int,
) {
	// 将tmpArray 中的元素从头到尾删除，直到tmpSum + num < target
	for {
		if tmpArray.Len() == 0 {
			break
		}

		tmpNum := tmpArray.Remove(tmpArray.Front()).(int)
		*tmpSum -= tmpNum

		if *tmpSum < target {
			break
		}
	}
}

// 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	if target <= 0 {
		return 0
	}

	if len(nums) == 0 {
		return 0
	}
	// 添加一个临时的数组，用来保存当前已经放入的元素
	tmpArray := list.New()

	tmpSum := 0

	// TODO: 添加变量，数组之和

	minLen := len(nums) + 1

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if num == target {
			return 1
		}

		// 如果临时元素为空，直接放入
		if tmpArray.Len() == 0 {
			tmpArray.PushBack(num)
			tmpSum = num
			continue
		}

		tmpArray.PushBack(num)
		tmpSum += num
		if tmpSum < target {
			continue
		}

		// XXXX: 最早出错的地方，这里没有考虑到，如果当前的数组之和已经大于等于target，那么就应该从头开始删除元素，直到数组之和小于target
		//if tmpArray.Len() < minLen {
		//	minLen = tmpArray.Len()
		//}

		fmt.Print("tmpSum: " + strconv.Itoa(tmpSum) + "\t")
		fmt.Print("num: " + strconv.Itoa(num) + "\t")
		fmt.Println("minLen: " + strconv.Itoa(minLen))
		utils.PrintList(tmpArray)

		for tmpSum >= target {
			// XXX: 写错的关键，这里没有考虑到可以动态比较长度，只考虑到了最后一个元素作为触发条件
			if tmpArray.Len() < minLen {
				minLen = tmpArray.Len()
			}

			tmpSum -= tmpArray.Remove(tmpArray.Front()).(int)
		}

	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}

func Perform(target int, nums []int) int {
	return minSubArrayLen(target, nums)
}
