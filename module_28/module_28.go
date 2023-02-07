package module_28

import (
	"bufio"
	"fmt"
	"os"
	stor "skillbox/module_28/storage"
	stud "skillbox/module_28/student"
	"strconv"
	"strings"
	"unicode"
)

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
