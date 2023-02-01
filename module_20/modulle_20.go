package module_20

import "log"

/*Задание 1
Напишите функцию, вычисляющую определитель матрицы размером 3 × 3.*/

func Determinant3x3(arr [][]int) int {
	return arr[0][0]*arr[1][1]*arr[2][2] +
		arr[0][1]*arr[1][2]*arr[2][0] +
		arr[0][2]*arr[1][0]*arr[2][1] -
		arr[0][2]*arr[1][1]*arr[2][0] -
		arr[0][0]*arr[1][2]*arr[2][1] -
		arr[0][1]*arr[1][0]*arr[2][2]
}

/*Задание 2
Напишите функцию, умножающую две матрицы размерами 3 × 5 и 5 × 4.*/

func MatrixMult(arr, brr [][]int) [][]int {
	result := make([][]int, len(arr))
	if len(arr[0]) != len(brr) {
		log.Fatal("Matrix multiplication not possible")
		return result
	}

	for i := range result {
		result[i] = make([]int, len(brr[0]))
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(brr[0]); j++ {
			result[i][j] = 0
			for k := 0; k < len(arr[0]); k++ {
				result[i][j] += arr[i][k] * brr[k][j]
			}
		}
	}
	return result
}
