package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

/*
Не знал нужно ли учитывать пробелы, на всякий случай учёл
*/
func main() {
	var input string
	var err error
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		println("Ошибка ввода", err)
		return
	}
	result := utf8.RuneCountInString(input)
	fmt.Println((result - 1))

}
