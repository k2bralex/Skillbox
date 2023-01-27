package function

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// CapitalWordsCount counts the number of words starting with a capital letter in a string
func CapitalWordsCount(str string) (count int) {
	slice := strings.Split(str, " ")
	for _, word := range slice {
		if unicode.IsUpper([]rune(word)[0]) {
			count++
		}
	}
	return
}

func Distinct(slice []int) (result []int) {
	k := make(map[int]bool)
	for _, v := range slice {
		if _, val := k[v]; !val {
			k[v] = true
			result = append(result, v)
		}
	}
	return
}

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

	return Distinct(digits)
}

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

func Ege() {

	var firstExam int
	var secondExam int
	var thirdExam int
	var totalScore int
	passingScore := 275

	fmt.Println("Введите результат первого экзамена:")
	fmt.Scan(&firstExam)

	fmt.Println("Введите результат второго экзамена:")
	fmt.Scan(&secondExam)

	fmt.Println("Введите результат третьего экзамена:")
	fmt.Scan(&thirdExam)

	totalScore = firstExam + secondExam + thirdExam

	fmt.Println("Сумма проходных баллов:", passingScore)
	fmt.Println("Количество набранных баллов:", totalScore)

	if totalScore >= passingScore {
		fmt.Println("Вы поступили.")
	} else {
		fmt.Println("Вы не поступили.")
	}
}

func GreaterThanFive() {
	var n1, n2, n3 int

	fmt.Println("Введите первое число:")
	fmt.Scan(&n1)

	fmt.Println("Введите второе число:")
	fmt.Scan(&n2)

	fmt.Println("Введите третье число:")
	fmt.Scan(&n3)

	if n1 > 5 || n2 > 5 || n3 > 5 {
		fmt.Println("Среди введённых чисел есть число больше 5.")
	} else {
		fmt.Println("Среди введённых чисел нет числа больше 5.")
	}
}

func ATM() {

	var withdrawal int
	fmt.Println("Введите сумму снятия со счёта:")
	fmt.Scan(&withdrawal)

	if withdrawal > 100000 || withdrawal%100 != 0 {
		fmt.Println("Операция не может быть выполнена")
	} else {
		fmt.Println("Операция успешно выполнена.\nВы сняли", withdrawal, "рублей.")
	}
}

func GreaterThanFiveCount() {
	count := 0

	for i := 0; i < 3; i++ {
		fmt.Println("Введите", i+1, "число:")
		var num int
		fmt.Scan(&num)
		if num >= 5 {
			count++
		}
	}

	fmt.Println("Среди введённых чисел", count, "больше или равны 5.")
}

func CheckCalculator() {
	var weekDay int
	var guests int
	var check float64

	fmt.Println("Введите день недели:")
	fmt.Scan(&weekDay)

	fmt.Println("Введите число гостей:")
	fmt.Scan(&guests)

	fmt.Println("Введите сумму чека:")
	fmt.Scan(&check)

	total := check

	if weekDay == 1 {
		total -= check * 0.1
		fmt.Println("Скидка по понедельникам:", check*0.1)
	}
	if weekDay == 5 && check >= 10000 {
		total -= check * 0.05
		fmt.Println("Скидка по пятницам:", check*0.05)
	}
	if guests > 5 {
		total += check * 0.1
		fmt.Println("Надбавка на обслуживание:", check*0.1)
	}
	fmt.Println("Сумма к оплате:", total)

}

func Students() {
	var totalStudents int
	var totalGroups int
	var id int

	fmt.Println("Всего студентов:")
	fmt.Scan(&totalStudents)

	fmt.Println("Всего групп:")
	fmt.Scan(&totalGroups)
	if totalGroups == 0 {
		panic("Деление на ноль")
	}

	fmt.Println("Порядковый номер студента:")
	fmt.Scan(&id)
	if id > totalStudents {
		panic("Id больше чем всего студентов")
	}

	group := id / (totalStudents / totalGroups)
	if id%(totalStudents/totalGroups) > 0 {
		group += 1
	}
	fmt.Println("Студент в группу:", group)
}
