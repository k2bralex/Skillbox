package module_9

import (
	"fmt"
	"math"
)

func OverflowCount() {
	var overflowUint8, overflowUint16 int
	var i uint8 = 0
	var k uint16 = 0

	for j := 0; j <= math.MaxUint32; j++ {
		i++
		k++
		if i == 0 {
			overflowUint8++
		}
		if k == 0 {
			overflowUint16++
		}
	}

	fmt.Println(overflowUint8, overflowUint16)
}

func MinType(a int16, b int16) string {
	mult := int(a) * int(b)

	switch {
	case (mult >= 0) && (mult <= int(math.MaxUint8)):
		return "uint8"
	case (mult >= 0) && (mult <= int(math.MaxUint16)):
		return "uint16"
	case (mult >= 0) && (mult <= int(math.MaxInt32)):
		return "uint32"
	case (mult >= int(math.MinInt8)):
		return "int8"
	case (mult >= int(math.MinInt16)):
		return "int16"
	case (mult >= int(math.MinInt32)):
		return "int32"
	}
	return "Unranged"
}
