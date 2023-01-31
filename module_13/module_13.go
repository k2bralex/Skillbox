package module_13

import (
	"fmt"
	"log"
)

/*Задание 1
Напишите функцию, которая принимает в качестве аргументов два числа типа int, вычисляет сумму чётных чисел
заданного диапазона и выводит результат в консоль. */

func Calc(x, y int, op func(int, int) int) {
	fmt.Println(op(x, y))
}

func Add(x, y int) int  { return x + y }
func Sub(x, y int) int  { return x - y }
func Mult(x, y int) int { return x * y }
func Div(x, y int) int {
	if y == 0 {
		log.Fatal("Division by zero")
		return 0
	}
	return x / y
}

/*Задание 2
Напишите функцию, которая принимает в качестве аргументов указатели на два типа int и меняет их значения местами.*/

func Swap(x, y *int) {
	*x, *y = *y, *x
}
