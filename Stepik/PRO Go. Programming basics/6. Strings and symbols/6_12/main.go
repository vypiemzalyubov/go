// Измените регистр символа, если он был латинской буквой: сделайте его заглавным, если он был строчной буквой и наоборот.
//
// Входные данные
// Задан единственный символ латинского алфавита c.
//
// Выходные данные
// Необходимо вывести получившийся символ.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	char, _, _ := reader.ReadRune()

	if char > 64 && char < 91 {
		fmt.Print(string(char + 32))
	} else {
		fmt.Print(string(char - 32))
	}
}
