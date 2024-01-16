package LCR013

import "fmt"

/*
给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回左上角 (row1, col1) 、右下角 (row2, col2) 的子矩阵的元素总和。
*/

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{matrix: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0

	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += this.matrix[i][j]

			fmt.Print(this.matrix[i][j])
			fmt.Print("; ")
		}
		fmt.Print("\n")
	}

	return sum
}
