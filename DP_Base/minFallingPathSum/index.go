package minFallingPathSum

import "math"

func minFallingPathSum(matrix [][]int) int {

	minVal := math.MaxInt32

	if len(matrix) == 0 {
		return 0
	}

	if len(matrix) == 1 {
		for i := 0; i < len(matrix[0]); i++ {
			minVal = min(minVal, matrix[0][i])
		}
		return minVal
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				matrix[i][j] = min(matrix[i-1][j], matrix[i-1][j+1]) + matrix[i][j]
			} else if j == len(matrix[i])-1 {
				matrix[i][j] = min(matrix[i-1][j-1], matrix[i-1][j]) + matrix[i][j]
			} else {
				matrix[i][j] = min(matrix[i-1][j-1], matrix[i-1][j], matrix[i-1][j+1]) + matrix[i][j]
			}
			if i == len(matrix)-1 {
				minVal = min(minVal, matrix[i][j])
			}
		}
	}

	return minVal
}
