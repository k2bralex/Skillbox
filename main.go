package main

import (
	fnc "Skillbox/function"
	"fmt"
)

func main() {

	/*	//Module 4
		fnc.Ege()
		fnc.GreaterThanFive()
		fnc.ATM()
		fnc.GreaterThanFiveCount()
		fnc.CheckCalculator()
		fnc.Students()

		//Module 5
		fnc.Qarter()
		fnc.IsAnyPositive()
		fnc.IsDistinct()
		fnc.NoChange()
		fnc.QuadraticEquationRoots()
		fnc.TicketCheck()
		fnc.GuessNumber()

		//Module 6
		fnc.IntSequence()
		fnc.SumByOne()
		fnc.Discount()
		fnc.LimitsCount()
		fnc.Apples()
		fnc.Elevator()

		//Module 7
		fnc.MirrorTicketCount()
		fnc.ChessBoard()
		fnc.FirTree()
		fnc.MaxDistance()

		//Module 8
		fnc.Seasons()
		fnc.WorkingDays()
		fmt.Println(fnc.GetChange([]int{5, 5, 5, 10, 20}))
		fmt.Println(fnc.GetChange([]int{10, 10}))
		fmt.Println(fnc.GetChange([]int{5, 5, 10, 10, 20}))

		//Module 9
		fnc.OverflowCount()
		fmt.Println(fnc.MinType(1, 1))
		fmt.Println(fnc.MinType(1, -1))
		fmt.Println(fnc.MinType(640, 100))
		fmt.Println(fnc.MinType(-640, 100))
		fmt.Println(fnc.MinType(3000, 3000))
		fmt.Println(fnc.MinType(-3000, 3000))

		//Module 10
		fnc.Tailor()
		fnc.Accountant()*/

	//Module 11
	exp := "Go is an Open source programming Language that makes it Easy to build simple, " +
		"reliable, and efficient Software"
	fmt.Println(fnc.CapitalWordsCount(exp))

	exp = "a10 10 20b 20 30c30 30 dd"
	fmt.Println(fnc.GetDigits(exp))
}
