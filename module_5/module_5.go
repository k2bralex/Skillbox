package module_5

import (
	"fmt"
	"math"
	"math/rand"
)

func Qarter() {
	var x, y int

	fmt.Println("Enter X, Y:")
	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		fmt.Println("1st qarter")
	} else if x < 0 && y > 0 {
		fmt.Println("2nd qarter")
	} else if x < 0 && y < 0 {
		fmt.Println("3rd qarter")
	} else if x > 0 && y < 0 {
		fmt.Println("4th qarter")
	} else {
		fmt.Println("origin")
	}
}

func IsAnyPositive() {
	for i := 0; i < 3; i++ {
		fmt.Println("Enter", i+1, "num:")
		var num int
		fmt.Scan(&num)
		if num < 0 {
			continue
		} else {
			fmt.Println("Entered number is positive")
			break
		}
	}
}

func IsDistinct() {
	var x, y, z int

	fmt.Println("Enter X, Y, Z:")
	fmt.Scan(&x, &y, &z)

	if x == y || y == z || z == x {
		fmt.Println("Numbers are'nt distinct")
	} else {
		fmt.Println("Numbers are distinct")
	}
}

func NoChange() {
	var cost int
	var coin1, coin2, coin3 int

	fmt.Println("Enter cost:")
	fmt.Scan(&cost)
	fmt.Println("Enter coins:")
	fmt.Scan(&coin1, &coin2, &coin3)

	if cost%(coin1+coin2+coin3) == 0 {
		fmt.Println("Is good. No change")
	} else {
		fmt.Println("No chance. Thare are change")
	}
}

func QuadraticEquationRoots() {
	var a, b, c float64

	fmt.Println("Enter a, b, c:")
	fmt.Scan(&a, &b, &c)

	var x1, x2 float64
	disc := math.Pow(b, 2) - 4*a*c

	if disc < 0 {
		fmt.Println("No roots.")
	} else {
		if disc == 0 {
			x1 = -b / (2 * a)
			x2 = x1
		} else {
			x1 = (-b + math.Sqrt(disc)) / (2 * a)
			x2 = (-b - math.Sqrt(disc)) / (2 * a)
		}
		fmt.Println("x1 =", x1, "x2 =", x2)
	}
}

func TicketCheck() {
	var ticket int

	fmt.Println("Enter number: ")
	fmt.Scan(&ticket)

	if isMirrorTicket(ticket) {
		fmt.Println("Ticket is a Mirror Ticket")
	} else if isLuckyTicket(ticket) {
		fmt.Println("Ticket is a Lucky Ticket")
	} else {
		fmt.Println("Ticket is a Simple Ticket")
	}

}

func isMirrorTicket(ticket int) bool {
	var arr [4]int

	for i := 0; i < 4; i++ {
		arr[i] = ticket % 10
		ticket /= 10
	}

	return (arr[0] == arr[3]) && (arr[1] == arr[2])
}

func isLuckyTicket(ticket int) bool {
	var arr [4]int

	for i := 0; i < 4; i++ {
		arr[i] = ticket % 10
		ticket /= 10
	}

	return arr[0]+arr[1] == arr[2]+arr[3]
}

func GuessNumber() {
	max := 10
	min := 1
	var guess int
	var answer string

	fmt.Println("Guess number from 1 to 10")

	for i := 0; i < 4; i++ {
		guess = rand.Intn(max-min) + min

		fmt.Println("Is your number is ", guess, "? y/n")
		fmt.Scan(&answer)

		if answer == "y" {
			fmt.Println("Cool")
			break
		} else {
			fmt.Println("Is number more than ", guess, "? y/n")
			fmt.Scan(&answer)
			if answer == "y" {
				min = guess
			} else {
				max = guess
			}
		}

	}
}
