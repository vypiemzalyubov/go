// Дано трехзначное число. Найдите сумму его цифр.
//
// Формат входных данных
// На вход дается трехзначное число.
//
// Формат выходных данных
// Выведите одно целое число — сумму цифр введенного числа.

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(n%10 + n/10%10 + n/100)
}