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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ = reader.ReadString('\n')
	result := utf8.RuneCountInString(input)
	fmt.Println((result - 1))

}
