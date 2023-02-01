package module_19

/*Задание 1
Напишите функцию, которая производит слияние двух отсортированных массивов длиной
четыре и пять в один массив длиной девять.*/

func ArrMerge(arr, brr []int) []int {
	i, j := 0, 0
	arrLen := len(arr)
	brrLen := len(brr)
	var crr []int

	for i < arrLen && j < brrLen {
		if arr[i] < brr[j] {
			crr = append(crr, arr[i])
			i++
		} else {
			crr = append(crr, brr[j])
			j++
		}
	}

	for i < arrLen {
		crr = append(crr, arr[i])
	}
	for j < brrLen {
		crr = append(crr, brr[i])
	}

	return crr
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
