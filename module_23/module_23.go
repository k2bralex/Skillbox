package module_23

/*
Задание 1. Сортировка вставками
Напишите функцию, сортирующую массив длины 10 вставками.
*/

func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		flag := 0
		for j := i - 1; j >= 0 && flag != 1; {
			if key < arr[j] {
				arr[j+1] = arr[j]
				j--
				arr[j+1] = key
			} else {
				flag = 1
			}
		}
	}
	return arr
}

/*Задание 2
Напишите анонимную функцию, которая на вход получает массив типа integer,
сортирует его пузырьком и переворачивает (либо сразу сортирует в обратном порядке, как посчитаете нужным).*/

func ReverseBubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
