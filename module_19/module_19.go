package module_19

import "sort"

/*Задание 1
Напишите функцию, которая производит слияние двух отсортированных массивов длиной
четыре и пять в один массив длиной девять.*/

func ArrMerge(arr, brr []int) []int {
	sort.Ints(arr)
	sort.Ints(brr)
	return append(arr, brr...)
}

/*Задание 2.
Отсортируйте массив длиной шесть пузырьком.*/

func BubbleSort(arr []int) []int {
	result := arr
	for i := 0; i <= len(arr)-2; i++ {
		for j := 0; j <= len(arr)-2; j++ {
			if result[j] > result[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return result
}
