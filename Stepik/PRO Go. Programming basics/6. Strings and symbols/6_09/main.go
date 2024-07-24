// Определите является ли данный символ цифрой или нет.
//
// Входные данные
// Задан единственный символ c.
//
// Выходные данные
// Необходимо вывести "YES", если символ является цифрой, и "NO" - в противном случае.

package main

import (
	"fmt"
)

func main() {
	var c rune
	_, _ = fmt.Scanf("%c", &c)
	flag := "NO"

	if c >= '0' && c <= '9' {
		flag = "YES"
	}
	fmt.Print(flag)
}
