package module_8

import "fmt"

func Seasons() {
	var seasons = map[string]string{
		"December":  "Winter",
		"January":   "Winter",
		"Febuary":   "Winter",
		"March":     "Spring",
		"April":     "Spring",
		"May":       "Spring",
		"June":      "Summer",
		"July":      "Summer",
		"August":    "Summer",
		"September": "Autumn",
		"October":   "Autumn",
		"November":  "Autumn",
	}
	fmt.Println("Enter month:")
	var value string
	fmt.Scan(&value)
	val, contains := seasons[value]
	if contains {
		fmt.Println(value, "is", val)
	} else {
		fmt.Println("Incorect input")
	}
}

func WorkingDays() {
	fmt.Print("Enter day: ")
	var value string
	fmt.Scan(&value)

	switch value {
	case "Mon":
		fmt.Println("Monday")
		fallthrough
	case "Tues":
		fmt.Println("Tuesday")
		fallthrough
	case "Wed":
		fmt.Println("Wednesday")
		fallthrough
	case "Thurs":
		fmt.Println("Thursday")
		fallthrough
	case "Fri":
		fmt.Println("Friday")
	default:
		fmt.Println("Non working day")
	}
}

func GetChange(bills []int) bool {
	five, ten := 0, 0

	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			ten++
			if five >= 1 {
				five--
			} else {
				return false
			}
		case 20:
			if five >= 3 {
				five -= 3
			} else if five >= 1 && ten >= 1 {
				five--
				ten--
			} else {
				return false
			}
		}
	}
	return true
}
