package LCR039

/*
*
请根据每日 气温 列表 temperatures ，重新生成一个列表，要求其对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
*/
func dailyTemperatures(temperatures []int) []int {

	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				result[i] = j - i
				break
			}
		}
	}

	return result
}
