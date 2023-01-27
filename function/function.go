package function

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// CapitalWordsCount counts the number of words starting with a capital letter in a string
func CapitalWordsCount(str string) (count int) {
	slice := strings.Split(str, " ")
	for _, word := range slice {
		if unicode.IsUpper([]rune(word)[0]) {
			count++
		}
	}
	return
}

func Distinct(slice []int) (result []int) {
	k := make(map[int]bool)
	for _, v := range slice {
		if _, val := k[v]; !val {
			k[v] = true
			result = append(result, v)
		}
	}
	return
}

func Tailor(x float64, eps float64) float64 {
	var res, term float64 = 1, 1
	eps = 1 / math.Pow(10, eps)
	for k := 1.0; term > eps; k++ {
		term = term * x / k
		res += term
	}
	return res
}

func GetDigits(str string) []int {
	var digits []int

	re := regexp.MustCompile("[0-9]+")
	current := re.FindAllString(str, -1)

	if len(current) < 1 {
		return digits
	}

	for _, val := range current {
		n, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		digits = append(digits, n)
	}

	return Distinct(digits)
}
