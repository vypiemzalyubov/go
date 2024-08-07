// По данной строке, найдите ее k-ый символ.
//
// Входные данные
// На вход программе дается непустая строка s, длины не больше 1000.
// На следующей строке вводится натуральное число k, не превосходящее 1000.
//
// Выходные данные
// Выведите k-ый символ строки, если он существует, в противном случае выведите "NO".

package main

import (
	"fmt"
)

func main() {
	var text string
	var k int
	fmt.Scan(&text, &k)

	if k <= len(text) {
		fmt.Print(string(text[k-1]))
	} else {
		fmt.Print("NO")
	}

}
