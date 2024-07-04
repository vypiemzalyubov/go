// Дана строка. Удалите k-ый символ в ней.
//
// Входные данные
// На первой строке вводится строка s. На второй строке вводится целое число 1≤k≤∣s∣, где ∣s∣ длина строки s.
//
// Выходные данные
// Необходимо вывести строку s без k-го символа.

package main

import (
	"fmt"
)

func main() {
	var text string
	var k int
	fmt.Scan(&text, &k)

	for i := 0; i < len(text); i++ {
		if i == k-1 {
			continue
		} else {
			fmt.Print(string(text[i]))
		}
	}
}
