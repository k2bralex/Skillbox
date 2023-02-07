package module_24

import (
	md14 "skillbox/module_14"
)

/*Задание 1
Напишите функцию, которая принимает массив чисел, а возвращает два массива: один из чётных чисел, второй из нечётных.*/

func SplitArray(arr []int) (brr, crr []int) {
	for _, v := range arr {
		if md14.IsEven(v) {
			brr = append(brr, v)
		} else {
			crr = append(crr, v)
		}
	}
	return
}

/*Задание 2
Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune,
а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово в
предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:

func parseTest(sentences []string, chars []rune) */

type Pair struct {
	Char  string
	Index int
}

func ParseTest(sentences []string, chars []rune) [][]Pair {
	var result [][]Pair
	for _, sentence := range sentences {
		var slice []Pair
		for i, r := range []rune(sentence) {
			if contains(chars, r) {
				slice = append(slice, Pair{Char: string(r), Index: i})
			}
		}
		result = append(result, slice)
	}

	return result
}

func contains(chars []rune, r rune) bool {
	for _, v := range chars {
		if v == r {
			return true
		}
	}
	return false
}
