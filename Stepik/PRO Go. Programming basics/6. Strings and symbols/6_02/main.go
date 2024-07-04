// Дана строка. Найдите перевернутую ей строку.
//
// Входные данные
// На вход подается строка, длина которой не превосходит 1000.
//
// Выходные данные
// Необходимо вывести ее перевертыш.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	text := scanner.Text()

	for i := len(text) - 1; i >= 0; i-- {
		fmt.Print(string(text[i]))
	}
}
