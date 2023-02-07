package module_25

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

/*Написать программу для нахождения подстроки в кириллической подстроке. Программа должна запускаться с помощью команды:
go run main.go --str "строка для поиска" --substr "поиска"*/

func General() {
	var str, sub string
	flag.StringVar(&str, "str", "", "find in")
	flag.StringVar(&sub, "substr", "", "find it")

	flag.Parse()

	val, err := isContains(&str, &sub)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)

}

func isContains(str, sub *string) (bool, error) {
	if *str == "" || *sub == "" {
		return false, fmt.Errorf("nil pointer")
	}
	return strings.Contains(*str, *sub), nil
}
