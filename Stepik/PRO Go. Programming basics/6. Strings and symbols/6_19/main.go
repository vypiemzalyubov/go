// Дан текст, составленный из букв латинского алфавита и цифр. Напечатать все имеющиеся в нем цифры.
//
// Входные данные
// На вход программе подается текст, составленный из букв латинского алфавита и цифр, длина которого не превосходит 1000.
//
// Выходные данные
// Выведите все имеющие цифры в тексте на одной строке через пробел.

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

	for i := 0; i < len(sentence); i++ {
		if sentence[i] > 47 && sentence[i] < 58 {
			fmt.Printf("%s ", string(sentence[i]))
		}
	}

}
