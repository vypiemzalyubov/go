// По данной строке определите, является ли она палиндромом? То есть, которое одинаково читается слева направо и справа налево. Например, слово "шалаш".
//
// Входные данные
// На вход подается одна строка без пробелов.
//
// Выходные данные
// Необходимо вывести  "YES", если строка является палиндромом, и "NO" - в противном случае.

package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scan(&text)
	flag := "YES"

	for i := 0; i < len(text); i++ {
		if string(text[i]) != string(text[len(text)-i-1]) {
			flag = "NO"
			break
		}
	}

	fmt.Print(flag)
}
