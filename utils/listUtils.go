package utils

func ArraySum(list []int) int {
	result := 0
	for _, val := range list {
		result += val
	}

	return result
}
