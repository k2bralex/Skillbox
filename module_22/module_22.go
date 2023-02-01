package module_22

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

/*Задание 1
Заполните массив неупорядоченными числами на основе генератора случайных чисел. Введите число.
Программа должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого.
При отсутствии введённого числа в массиве — вывести 0. Для удобства проверки реализуйте вывод массива на экран.*/

// JumpSearch Implement Jump Search algorithm to find number in array.
// On each iteration jump on sqrt(arraylength) while current value smaller than searched number
// Than use linear search between last and current step
func JumpSearch(size, num int) (int, int) {
	arr := FillArray(size)
	sort.Ints(arr)

	jumpStep := int(math.Sqrt(float64(size)))
	lastStep := 0

	for arr[int(math.Min(float64(jumpStep), float64(size-1)))] < num {
		lastStep = jumpStep
		jumpStep += int(math.Sqrt(float64(size)))
		if lastStep >= size {
			return -1, 0
		}
	}

	for arr[lastStep] < num {
		lastStep++
		if lastStep == int(math.Min(float64(jumpStep), float64(size))) {
			return -1, 0
		}
	}

	if arr[lastStep] == num {
		return lastStep, len(arr[lastStep+1:])
	}

	return 0, 0
}

func FillArray(size int) []int {
	rand.NewSource(time.Now().UnixNano())
	var result []int
	for i := 0; i < size; i++ {
		result = append(result, rand.Intn(200)-100)
	}
	return result
}

/*Задание 2
Заполните упорядоченный массив из 12 элементов и введите число. Необходимо реализовать поиск первого
вхождения заданного числа в массив. Сложность алгоритма должна быть минимальная.*/

// FindFirstIn implement binary search for finding the first occurrence
func FindFirstIn(num int, arr []int) int {
	min := 0
	max := len(arr) - 1
	for min <= max {
		mid := (max-min)/2 + min
		if arr[mid] == num && (mid == 0 || arr[mid-1] < num) {
			return mid
		} else if arr[mid] < num {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}
