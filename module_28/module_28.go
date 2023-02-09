package module_28

import (
	"bufio"
	"fmt"
	"os"
	i "skillbox/module_28/instance"
	stor "skillbox/module_28/storage"
	stud "skillbox/module_28/student"
	t "skillbox/module_28/teacher"
	"strconv"
	"strings"
	"unicode"
)

func RunWithGoroutines() {

	userInputChan := reader()

}

func reader() *chan string {
	out := make(chan string)
	var input string
	go func() {
		for input != "exit" {
			out <- userInput()
		}
		close(out)
	}()
	return &out
}

func creator(in *chan string) *chan i.Worker {
	out := make(chan i.Worker)
	go func() {
		for v := range out {

		}
	}()
	return &out
}

func Run() {
	teachers := [...]t.Teacher{
		{"Teacher1", 30, 20000.99},
		{"Teacher2", 31, 30000.99},
		{"Teacher3", 32, 40000.99},
		{"Teacher4", 33, 50000.99},
	}

	students := [...]stud.Student{
		{"Student1", 18, 22},
		{"Student2", 17, 21},
		{"Student3", 19, 20},
		{"Student4", 19, 19},
	}

	group := stor.Group{}
	for i := 0; i < len(teachers)*2; i++ {
		if i < 4 {
			err := group.Put(&teachers[i])
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			err := group.Put(&students[i-len(students)])
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	group.PrintGroup()
}

func GroupCreate() {
	group := stor.Group{}
	for {
		fmt.Print(" -> ")
		input := userInput()
		if input == "exit" {
			group.PrintGroup()
			return
		}

		details := strings.Fields(input)
		if len(details) < 3 {
			fmt.Println("Not enough arguments")
			continue
		}

		name, err0 := nameCheck(details[0])
		if err0 != nil {
			fmt.Println(err0.Error())
			continue
		}
		age, err1 := strconv.ParseInt(details[1], 10, 16)
		if err1 != nil {
			fmt.Println(err1.Error())
			continue
		}
		grade, err2 := strconv.ParseInt(details[2], 10, 16)
		if err2 != nil {
			fmt.Println(err2.Error())
			continue
		}

		if err := group.Put(stud.NewStudent(name, int(age), int(grade))); err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}

func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return scanner.Text()
}

func nameCheck(str string) (string, error) {
	for _, v := range []rune(str) {
		if !unicode.IsLetter(v) {
			return "", fmt.Errorf("name non alphabetic")
		}
	}
	return str, nil
}
