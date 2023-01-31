package main

import (
	/*	md4 "skillbox/module_4"
		md5 "skillbox/module_5"
		md6 "skillbox/module_6"
		md7 "skillbox/module_7"
		md8 "skillbox/module_8"
		md9 "skillbox/module_9"
		md10 "skillbox/module_10"
		md11 "skillbox/module_11"
		md12 "skillbox/module_12"
		md13 "skillbox/module_13"
		md14 "skillbox/module_14"	*/

	"fmt"
	md15 "skillbox/module_15"
)

func main() {

	//Module 15
	e, o := md15.EvenOddCount()
	fmt.Printf("Evens: %d\nOdds: %d\n", e, o)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%v -> ", arr)
	md15.Reverse(arr)
	fmt.Printf("%v\n", arr)

	/*
		//Module 14

		fmt.Println(md14.IsEven(10))
		fmt.Println(md14.ChangeCoords())
		fmt.Println(md14.NumConvert(2))
		fmt.Println(md14.GlobalVarCalc1(5))

		//Module 13

		md13.Calc(1, 2, md13.Add)
		md13.Calc(4, 2, md13.Sub)
		md13.Calc(1, 6, md13.Mult)
		md13.Calc(6, 3, md13.Div)

		a, b := 3, 4
		fmt.Printf("%d, %d -> ", a, b)
		md13.Swap(&a, &b)
		fmt.Printf("%d, %d\n", a, b)

		//Module 12
		md12.AnnouncementWriter()
		md12.PermissionMode("read_only")
		fmt.Println(md12.ParenthesesGenerator(3))

		//Module 11
		exp := "Go is an Open source programming Language that makes it Easy to build simple, " +
			"reliable, and efficient Software"
		fmt.Println(md11.CapitalWordsCount(exp))

		exp = "a10 10 20b 20 30c30 30 dd"
		fmt.Println(md11.GetDigits(exp))

		/*
		//Module 10
		md10.Tailor()
		md10.Accountant()

		//Module 9
		md9.OverflowCount()
		fmt.Println(md9.MinType(1, 1))
		fmt.Println(md9.MinType(1, -1))
		fmt.Println(md9.MinType(640, 100))
		fmt.Println(md9.MinType(-640, 100))
		fmt.Println(md9.MinType(3000, 3000))
		fmt.Println(md9.MinType(-3000, 3000))

		//Module 8
		md8.Seasons()
		md8.WorkingDays()
		fmt.Println(md8.GetChange([]int{5, 5, 5, 10, 20}))
		fmt.Println(md8.GetChange([]int{10, 10}))
		fmt.Println(md8.GetChange([]int{5, 5, 10, 10, 20}))

		//Module 7
		md7.MirrorTicketCount()
		md7.ChessBoard()
		md7.FirTree()
		md7.MaxDistance()

		//Module 6
		md6.IntSequence()
		md6.SumByOne()
		md6.Discount()
		md6.LimitsCount()
		md6.Apples()
		md6.Elevator()

		//Module 5
		md5.Qarter()
		md5.IsAnyPositive()
		md5.IsDistinct()
		md5.NoChange()
		md5.QuadraticEquationRoots()
		md5.TicketCheck()
		md5.GuessNumber()

		//Module 4
		md4.Ege()
		md4.GreaterThanFive()
		md4.ATM()
		md4.GreaterThanFiveCount()
		md4.CheckCalculator()
		md4.Students()*/
}
