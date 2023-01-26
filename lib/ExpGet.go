package lib

import (
	"math"
)

func Tailor(x float64, eps float64) float64 {
	var res, term float64 = 1, 1
	eps = 1 / math.Pow(10, eps)
	for k := 1.0; term > eps; k++ {
		term = term * x / k
		res += term
	}
	return res
}
