// Дано предложение, составленное из строчных букв латинского алфавита. Определить, сколько в нем гласных букв. Гласными считаются буквы  a, e, i, o, u.
//
// Входные данные
// На вход программе подается строка, длина которого не превосходит 1000.
//
// Выходные данные
// Выведите количество гласных букв.

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
	cnt := 0

	for i := 0; i < len(sentence); i++ {
		if string(sentence[i]) == "a" || string(sentence[i]) == "e" || string(sentence[i]) == "i" || string(sentence[i]) == "o" || string(sentence[i]) == "u" {
			cnt++
		}
	}

	fmt.Print(cnt)
}
