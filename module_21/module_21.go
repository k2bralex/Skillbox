package module_21

import (
	"fmt"
	"log"
)

/*Задание 1
Напишите функцию, производящую следующие вычисления.
S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.
Тип S должен быть во float32.*/

func Upcast(x int16, y uint8, z float32) float32 {
	if z == 0 {
		log.Fatal("Division by zero")
		return 0
	}
	return 2*float32(x) + float32(y)*float32(y) - 3/z
}

/*Задание 2
Напишите функцию, которая на вход принимает функцию вида A func (int, int) int, а внутри оборачивает
и вызывает её при выходе (через defer).
Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций могут быть любыми, но главное,
чтобы все три выполняли разное действие.*/

func UsingDefer(a1, a2 int, action func(int, int) int) {
	defer fmt.Println(action(a1, a2))
}

var F1 = func(x, y int) int { return x + y }
var F2 = func(x, y int) int { return x - y }
var F3 = func(x, y int) int { return x * y }
