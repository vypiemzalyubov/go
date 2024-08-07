// По заданной строчной букве латинского алфавита, выведите все строчные буквы латинского алфавита, начиная от начала до заданной буквы.
//
// Входные данные
// Задан единственный символ c - строчная буква латинского алфавита.
//
// Выходные данные
// Выведите на одной строке и через пробел все строчные буквы латинского алфавита, начиная от начала и заканчивая символом c.

package main

import (
	"fmt"
)

func main() {
	var c rune
	_, _ = fmt.Scanf("%c", &c)

	for w := 'a'; w <= c; w++ {
		fmt.Printf("%s ", string(w))
	}
}
