package module_10

import (
	"fmt"
	"math"
)

func Tailor() {

	var x, eps float64
	var res, term float64 = 1, 1

	fmt.Println("Enter num and accuracy:")
	fmt.Scan(&x, &eps)

	eps = 1 / math.Pow(10, eps)

	for k := 1.0; term > eps; k++ {
		term = term * x / k
		res += term
	}

	fmt.Println("Answer: ", res)
}

func Accountant() {
	var deposit, persent float64
	var years int
	var bankChange float64 = 0

	fmt.Println("Enter deposit, persentage and years:")
	fmt.Scan(&deposit, &persent, &years)

	for i := 0; i < years*12; i++ {
		monthlyIncrease := deposit * persent / 100
		deposit += math.Floor(monthlyIncrease*100) / 100
		bankChange += monthlyIncrease - math.Floor(monthlyIncrease*100)/100
	}

	fmt.Println("Man: ", deposit, "Bank: ", bankChange)
}
