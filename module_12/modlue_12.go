package module_12

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

/*Задание 1
Напишите программу, которая на вход получала бы строку, введённую пользователем, а в файл писала № строки,
дату и сообщение в формате:	2020-02-10 15:00:00 продам гараж.
При вводе слова exit программа завершает работу.*/

func AnnouncementWriter() {
	fileName := "announcement.txt"
	//fileName := "empty_file.txt"

	//Open file if exits. If no - create
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	fileSize, err := file.Stat()
	if err != nil {
		log.Fatal(err)
		return
	}

	var text string
	var strCount int

	//read and count lines if anything in file
	if fileSize.Size() > 0 {
		text, strCount = readFile(fileName, int(fileSize.Size()))
		fmt.Println(text)
	} else {
		fmt.Println("File is empty")
	}

	//ask for new string from user and append formatted string into file
	for {
		fmt.Print("Write new announcement:")
		input := customConsoleReader()
		if input == "exit\n" || input == "Exit\n" {
			fmt.Println(input)
			break
		}
		strCount++
		fmt.Fprint(file, strconv.Itoa(strCount)+" "+time.DateOnly+" "+input)
	}

}

// customConsoleReader Read console input while "Return" entered
func customConsoleReader() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return input
}

// readFile read wile and count strings
func readFile(fileName string, fileSize int) (string, int) {
	stringInFile := 0
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return "", 0
	}

	data := make([]byte, fileSize)
	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
	}
	for _, val := range data {
		if val == '\n' {
			stringInFile++
		}
	}
	file.Close()
	return string(data), stringInFile
}

/*Задание 2
Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике,
без использования ioutil. Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.*/

//Done in task 1

/*Задание 3.
Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.*/

// PermissionMode check if file exist - delete and create new one.
// Set permission mode
// Try to write into file
func PermissionMode(name string) {
	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		os.Remove(name)
	}

	file, err := os.Create(name)
	if err := os.Chmod(name, 0444); err != nil {
		log.Fatal(err)
		return
	}
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString("Try write something to file\n"); err != nil {
		fmt.Println(err.Error())
		return
	}
}

/*Задание 4.
Перепишите задачи 1 и 2, используя пакет ioutil.*/

// ioutil package is deprecated

/*Задание 5 (по желанию)
Напишите программу, которая на вход принимала бы интовое число и для него генерировала
бы все возможные комбинации круглых скобок.*/

// ParenthesesGenerator sent into recursive function link on final strings slice
func ParenthesesGenerator(num int) (result []string) {
	generate(&result, "", 0, 0, num)
	return
}

// get pointer of slice and append new string into final slice
func generate(result *[]string, str string, open int, close int, i int) {
	if open+close == 2*i {
		*result = append(*result, str)
		return
	}
	if open < i {
		generate(result, str+"(", open+1, close, i)
	}
	if close < open {
		generate(result, str+")", open, close+1, i)
	}

}
