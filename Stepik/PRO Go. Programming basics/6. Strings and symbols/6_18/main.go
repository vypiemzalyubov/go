// Дано предложение, составленное из строчных букв латинского алфавита. Все буквы 'e' в нем заменить буквой 'i'.
//
// Входные данные
// На вход программе подается строка, длина которого не превосходит 1000.
//
// Выходные данные
// Выведите строку, в котором все буквы 'e' заменены буквой 'i'.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sentence := scanner.Text()

	for _, char := range sentence {
		if string(char) == "e" {
			fmt.Print("i")
		} else {
			fmt.Print(string(char))
		}
	}
}
