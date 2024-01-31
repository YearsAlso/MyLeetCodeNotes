package LCR039

/*
给定非负整数数组 heights ，数组中的数字用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)

	result := 0

	// 使用单调栈
	// 1. 如果当前元素大于栈顶元素，直接入栈
	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 || heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			// 2. 如果当前元素小于栈顶元素，弹出栈顶元素，计算面积
			// 3. 如果栈为空，说明当前元素是最小元素，宽度为i，否则宽度为i-stack[len(stack)-1]-1
			// 4. 面积为宽度乘以高度
			for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
				height := heights[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
				width := i
				if len(stack) > 0 {
					width = i - stack[len(stack)-1] - 1
				}
				area := height * width
				if area > result {
					result = area
				}
			}

			// 5. 当前元素入栈,这一步容易忘记
			stack = append(stack, i)
		}
	}

	// 6. 如果栈不为空，重复步骤2-5
	for len(stack) > 0 {
		height := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		width := len(heights)
		if len(stack) > 0 {
			width = len(heights) - stack[len(stack)-1] - 1
		}
		area := height * width
		if area > result {
			result = area
		}
	}

	return result
}
