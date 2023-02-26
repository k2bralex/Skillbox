package module_28

import (
	"bufio"
	"fmt"
	"os"
	stor "skillbox/module_28/storage"
	stud "skillbox/module_28/student"
	t "skillbox/module_28/teacher"
	"strconv"
	"strings"
	"unicode"
)

// RunWithGoroutines use goroutines to get internal input from model and create entities for storage
func RunWithGoroutines() {
	in := make(chan string)
	g := stor.Group{}

	go reader(in)

	for {
		v, ok := <-in
		if !ok {
			g.PrintGroup()
			break
		}
		err := g.Put(creator(v))
		if err != nil {
			continue
		}
	}
}

func reader(in chan string) {
	go func() {
		for {
			input := userInput()
			if input == "exit" {
				break
			}
			in <- input
		}
		close(in)
	}()
}

func creator(in string) *stud.Student {
	n, a, g, e := validate(in)
	if e != nil {
		return nil
	}
	return &stud.Student{Name: n, Age: int(a), Grade: int(g)}
}

/*-------------------------------------------------------------------------------------------*/

// Run use interface to get any service into storage
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

/*-------------------------------------------------------------------------------------------*/

// GroupCreate base module task create storage, get input from model, validate due to service model,
// put into storage, print storage after input
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

func validate(input string) (string, int64, int64, error) {
	details := strings.Fields(input)
	if len(details) < 3 {
		return "", 0, 0, fmt.Errorf("not enough arguments")
	}

	name, err0 := nameCheck(details[0])
	if err0 != nil {
		return "", 0, 0, err0
	}
	age, err1 := strconv.ParseInt(details[1], 10, 0)
	if err1 != nil {
		return "", 0, 0, err1
	}
	grade, err2 := strconv.ParseInt(details[2], 10, 0)
	if err2 != nil {
		return "", 0, 0, err2
	}
	return name, age, grade, nil
}

func nameCheck(str string) (string, error) {
	for _, v := range []rune(str) {
		if !unicode.IsLetter(v) {
			return "", fmt.Errorf("name non alphabetic")
		}
	}
	return str, nil
}
