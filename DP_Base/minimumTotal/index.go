package minimumTotal

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	minVal := math.MaxInt
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j >= len(triangle[i-1]) {
				triangle[i][j] += triangle[i-1][j-1]
			} else if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else {
				triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
			}

			if i == len(triangle)-1 {
				if triangle[i][j] < minVal {
					minVal = triangle[i][j]
				}
			}
		}
	}

	return minVal
}
