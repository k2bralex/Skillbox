package module_7

import "fmt"

func MirrorTicketCount() {
	var count int

	for i := 100000; i <= 999999; i++ {
		if i%1000 == i/1000 {
			count++
		}
	}

	fmt.Println(count)
}

func ChessBoard() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			if (i+j)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func FirTree() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= 5-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("*")
		}
		for t := 1; t <= i; t++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func MaxDistance() {
	var numbers []int
	maxDistance := 0
	for i := 100000; i <= 999999; i++ {
		if isLucky(i) {
			numbers = append(numbers, i)
			var currentDistance int
			if len(numbers) > 2 {
				currentDistance = numbers[len(numbers)-1] - numbers[len(numbers)-2]
			}
			if maxDistance < currentDistance {
				maxDistance = currentDistance
			}
		}
	}
	fmt.Println(maxDistance)
}

func isLucky(num int) bool {
	var arr [6]int
	for i := 0; i < 6; i++ {
		arr[i] = num % 10
		num /= 10
	}
	return arr[0]+arr[1]+arr[2] == arr[3]+arr[4]+arr[5]
}
