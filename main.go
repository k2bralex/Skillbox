package main

import (
	fnc "Skillbox/function"
	"fmt"
)

func main() {
	//Task 1
	exp := "Go is an Open source programming Language that makes it Easy to build simple, " +
		"reliable, and efficient Software"
	fmt.Println(fnc.CapitalWordsCount(exp))

	//Task 2
	exp = "a10 10 20b 20 30c30 30 dd"
	fmt.Println(fnc.GetDigits(exp))
}
