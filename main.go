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
		md14 "skillbox/module_14"
		md15 "skillbox/module_15"
		md19 "skillbox/module_19"
		md20 "skillbox/module_20"
		md21 "skillbox/module_21"
		md20 "skillbox/module_20"
		md21 "skillbox/module_21"
		md22 "skillbox/module_22"
		md23 "skillbox/module_23"
		md24 "skillbox/module_24"
		md25 "skillbox/module_25"	*/

	//"log"
	//"reflect"
	//"os"
	"os"
	md26 "skillbox/module_26"
)

func main() {

	//Module 26

	md26.ConcatTextFile(os.Args[1:])

	/*


		//Module 25

		if len(os.Args[1:]) < 4 {
			log.Fatal("Not enough arguments")
		}
		md25.General()

		//Module 24

		arr := []int{13, 86, 60, 46, 73, 52, 73, 57, 49, 99, 133, 20, 1}
		brr, crr := md24.SplitArray(arr)
		fmt.Printf("Even: %v\nOdd:  %v\n", brr, crr)

		sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
		chars := []rune{'H', 'E', 'L', 'П', 'М'}

		fmt.Println(md24.ParseTest(sentences, chars))


		//Module 23

		arr := []int{13, 86, 60, 46, 73, 52, 73, 57, 49, 99, 133, 20, 1}
		fmt.Println(md23.InsertSort(arr))

		sort := func(arr []int) []int {
			for i := 0; i < len(arr)-1; i++ {
				for j := 0; j < len(arr)-i-1; j++ {
					if arr[j] < arr[j+1] {
						arr[j], arr[j+1] = arr[j+1], arr[j]
					}
				}
			}
			return arr
		}
		fmt.Println(sort(arr))
		//fmt.Println(md23.ReverseBubbleSort(arr))

		//Module 22

		a, b := md22.JumpSearch(1000, 10)
		fmt.Println(a, b)

		num := 3
		fmt.Printf("Number %d. First in on index: %d\n", num,
			md22.FindFirstIn(num, []int{1, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5, 6, 7, 7, 7, 8, 9, 10}))

		//Module 21

		var (
			x int16   = -3555
			y uint8   = 3
			z float32 = 6.85
		)
		res := md21.Upcast(x, y, z)
		fmt.Printf("Result: %.2f, value type: %s\n", res, reflect.TypeOf(res))

		md21.UsingDefer(10, 25, md21.F1)
		md21.UsingDefer(10, 25, md21.F2)
		md21.UsingDefer(10, 25, md21.F3)

		//Module 20

		fmt.Println(md20.Determinant3x3([][]int{
			{2, 4, 6},
			{0, 9, 7},
			{3, 1, 0},
		}))

		arr := [][]int{
			{2, 4, 6, 9, 11},
			{0, -4, 5, 3, 1},
			{1, 7, 3, -24, 0},
		}

		brr := [][]int{
			{4, 3, 11, 9},
			{0, -4, 5, 3},
			{2, 17, -3, -24},
			{4, 7, 1, -24},
			{5, -7, 0, 10},
		}

		fmt.Println(md20.MatrixMult(arr, brr))

		//Module 19
		fmt.Printf("%v + %v -> %v\n",
			[]int{2, 3, 1, 0},
			[]int{9, 7, 6, 8, 5},
			md19.ArrMerge([]int{2, 3, 1, 0}, []int{9, 7, 6, 8, 5}))

		arr := []int{8, 9, 1, 0, -6, -88}
		fmt.Printf("%v -> %v\n", arr, md19.BubbleSort(arr))

		//Module 15
		e, o := md15.EvenOddCount()
		fmt.Printf("Evens: %d\nOdds: %d\n", e, o)

		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		fmt.Printf("%v -> %v\n", arr, md15.Reverse(arr))

		fmt.Printf("%v -> ", arr)
		md15.SelfReverse(arr)
		fmt.Printf("%v\n", arr)

		//Module 14

		fmt.Println(md14.IsEven(10))
		fmt.Println(md14.ChangeCoords())
		fmt.Println(md14.NumConvert(2))
		fmt.Println(md14.GlobalVarCalc1(GlobalVarCalc2(GlobalVarCalc3(5))))

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
