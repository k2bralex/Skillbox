package module_4

import "fmt"

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
