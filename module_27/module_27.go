package module_27

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*Напишите программу, которая считывает ввод с stdin, создаёт структуру student и записывает
указатель на структуру в хранилище map[studentName] *Student.
Программа должна получать строки в бесконечном цикле, создать структуру Student через функцию newStudent,
далее сохранить указатель на эту структуру в map, а после получения EOF (ctrl + d) вывести на экран
имена всех студентов из хранилища. Также необходимо реализовать методы put, get. */

// Student is a structure for mapping
type Student struct {
	name       string
	age, grade int
}

// Conditional constructor
func newStudent(name string, age, grade int) *Student {
	return &Student{name: name, age: age, grade: grade}
}

// GroupCreate creates map of students
//get user input while no "exit"
//validate data
//create new instance
//add to map
//print results

func GroupCreate() {
	group := map[string]*Student{}

	for {
		fmt.Print(" -> ")
		input := userInput()
		if input == "exit" {
			printMap(&group)
			return
		}

		details := strings.Fields(input)
		if len(details) < 3 {
			fmt.Println("Not enough arguments")
			continue
		}

		name := details[0]
		age, err1 := strconv.ParseInt(details[1], 10, 16)
		grade, err2 := strconv.ParseInt(details[2], 10, 16)
		if err1 != nil || err2 != nil {
			fmt.Println("Incorrect arguments")
			continue
		}

		if _, ok := group[name]; ok {
			fmt.Println("Account exist")
			continue
		}

		group[name] = newStudent(name, int(age), int(grade))

	}
}

func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

func printMap(m *map[string]*Student) {
	for k, v := range *m {
		fmt.Println(" <-", k, v.age, v.grade)
	}
}
