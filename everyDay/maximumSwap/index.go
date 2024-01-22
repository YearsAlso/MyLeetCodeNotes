package maximumSwap

// 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

func maximumSwap(num int) int {
	result := 0

	if num < 10 {
		return num
	}

	// 先找到最高位

	bitList := make([]int, 0)
	srcNum := num
	// TODO: 不一定是最高位，也可能是比自己小的最高位
	maxNumber := 0
	maxNumberIndex := 0

	for i := 0; num > 0; i++ {
		elems := num % 10
		bitList = append(bitList, elems)
		if elems > maxNumber {
			maxNumber = elems
			maxNumberIndex = i
		}
		num /= 10
	}

	// 从最高位开始，找到第一个比最高位小的数，交换
	needSwap := false

	for i := len(bitList) - 1; i >= 1; i-- {
		// i 越大，位数越高
		if bitList[i] < maxNumber && i > maxNumberIndex {
			// 交换
			bitList[i], bitList[maxNumberIndex] = bitList[maxNumberIndex], bitList[i]
			needSwap = true
			break
		}
	}

	if !needSwap {
		return srcNum
	}

	// 重新组合
	for i := len(bitList) - 1; i >= 0; i-- {
		result = result*10 + bitList[i]
	}

	return result
}
