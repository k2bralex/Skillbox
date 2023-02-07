package module_26

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*Написать программу аналог cat.
Программа должна получать на вход имена двух файлов, необходимо  конкатенировать их содержимое, используя strings.Join.
При получении одного файла на входе программа должна печатать его содержимое на экран.
При получении двух файлов на входе программа соединяет их и печатает содержимое обоих файлов на экран.
Если программа запущена командой go run firstFile.txt secondFile.txt resultFile.txt, то она должна написать
два соединённых файла в результирующий.*/

// ConcatTextFile get arguments from command line
// Due to len of args print files or combine into new one
func ConcatTextFile(args []string) {
	switch len(args) {
	case 1:
		fmt.Printf("Open file: %s\n%s", args[0], fileReader(args[0]))
	case 2:
		fmt.Printf("Open file: %s\n%s\n\nOpen file: %s\n%s",
			args[0], fileReader(args[0]), args[1], fileReader(args[1]))
	case 3:
		fileCombiner(args)
	}
}

// Open file, read content, close, return content as string
func fileReader(name string) string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	return string(text)
}

// Open files, read content, merge content, write into new file
func fileCombiner(args []string) {
	text1 := fileReader(args[0])
	text2 := fileReader(args[1])

	result := strings.Join([]string{text1, text2}, "\n")

	file, err := os.Create(args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(result)
}
