package module_14

import (
	"math/rand"
	"time"
)

/*Задание 1
Напишите функцию, которая на вход получает число и возвращает true, если число четное, и false, если нечётное.*/

func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

/*Задание 2
Напишите программу, которая с помощью функции генерирует три случайные точки в двумерном пространстве
(две координаты), а затем с помощью другой функции преобразует эти координаты по формулам:
	x = 2 × x + 10, y = −3 × y − 5.*/

func ChangeCoords() (int, int) {
	x, y := getRandomPoint()
	return 2*x + 10, -3*y - 5
}

func getRandomPoint() (int, int) {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(200) - 100, rand.Intn(200) - 100
}

/*Задание 3
Напишите программу, которая на вход получает число, затем с помощью двух функций преобразует его.
Первая умножает, а вторая прибавляет число, используя именованные возвращаемые значения.*/

func NumConvert(num int) (res int) {
	res = add(mult(num))
	return
}

func mult(num int) (res int) {
	num *= 2
	return
}

func add(num int) (res int) {
	res += 2
	return
}

/*Задание 4
Напишите программу, в которой будет три функции, попарно использующие разные глобальные переменные.
Функции должны прибавлять к поданному на вход числу глобальную переменную и возвращать результат.
Затем вызовите по очереди три функции, передавая результат из одной в другую.*/

var NUM = 10

func GlobalVarCalc1(num int) int {
	return GlobalVarCalc2(num) + NUM
}

func GlobalVarCalc2(num int) int {
	return GlobalVarCalc3(num) + NUM
}

func GlobalVarCalc3(num int) int {
	return num + NUM
}
