package module_6

import (
	md11 "Skillbox/module_11"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func IntSequence() {
	var num int

	fmt.Println("Enter number:")
	fmt.Scan(&num)

	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
}

func SumByOne() {
	var n1, n2 int
	sum := n1

	fmt.Println("Enter number:")
	fmt.Scan(&n1, &n2)

	for sum != (n1 + n2) {
		sum += 1
	}
	fmt.Println(sum)
}

func Discount() {
	var price, discount float64

	fmt.Println("Enter price and discount:")
	fmt.Scan(&price, &discount)

	if discount > 30 || price*discount/100 > 2000 {
		fmt.Println("Discount too hight")
	} else {
		fmt.Println("Fianl price:", price*(1-discount/100))
	}
}

func LimitsCount() {
	a := 0
	b := 0
	c := 0
	for {
		if a != 10 {
			a++
		} else if b != 100 {
			b++
		} else if c != 1000 {
			c++
		} else {
			break
		}
	}

	fmt.Println(a, b, c)
}

func Apples() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Enter", i, "basket volume:")
		var volume int
		fmt.Scan(&volume)
		for j := 1; j <= volume; j++ {
			fmt.Println(j, "apples in basket", i)
		}
	}
}

func Elevator() {
	passengerInElevator := 0
	floors := []int{0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := len(floors); i > 0; i-- {

		fmt.Println("Elevator on ", i, " floor")

		if floors[i-1] > 0 && passengerInElevator < 2 {
			floors[i-1]--
			passengerInElevator++
			fmt.Println(passengerInElevator, "passenger inside")
			i++
			continue
		}
		if i == 1 {
			fmt.Println(passengerInElevator, "Passengers get out of elevator\nElevator goes upstairs")
		}
		if i == 1 && Sum(floors) > 0 {
			i = len(floors) + 1
			passengerInElevator = 0
		}
	}

	fmt.Println("Passengers awating\n", floors)
}

func Sum(array []int) int {
	arrSum := 0
	for _, a := range array {
		arrSum = arrSum + a
	}
	return arrSum
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

	return md11.Distinct(digits)
}
