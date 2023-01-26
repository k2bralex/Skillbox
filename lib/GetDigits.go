package lib

import (
	"log"
	"regexp"
	"strconv"
)

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
