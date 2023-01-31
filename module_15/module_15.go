package module_15

import (
	"fmt"
	"log"
	md14 "skillbox/module_14"
)

/*Задание 1
Разработайте программу, позволяющую ввести 10 целых чисел, а затем вывести из них количество
чётных и нечётных чисел. Для ввода и подсчёта используйте разные циклы.*/

func EvenOddCount() (even int, odd int) {
	slice := InputGet()
	for i := range slice {
		if md14.IsEven(slice[i]) {
			even++
		} else {
			odd++
		}
	}
	return even, odd
}

func InputGet() []int {
	slice := make([]int, 10)
	fmt.Print("Enter 10 integers: ")
	for i := range slice {
		_, err := fmt.Scan(&slice[i])
		if err != nil {
			log.Fatal(err)
			return slice
		}
	}
	return slice
}

/*Задание 2
Напишите функцию, принимающую на вход массив и возвращающую массив, в котором элементы идут в обратном порядке
по сравнению с исходным. Напишите программу, демонстрирующую работу этого метода.*/

func Reverse(arr []int) []int {
	l := len(arr)
	reversed := make([]int, l)
	for i := 0; i < l/2; i++ {
		reversed[i], reversed[l-1-i] = arr[l-1-i], arr[i]
	}
	return reversed
}

func SelfReverse(arr []int) {
	l := len(arr)
	var head, tail *int
	for i := 0; i < l/2; i++ {
		head, tail = &arr[i], &arr[l-1-i]
		*head, *tail = *tail, *head
	}
}
