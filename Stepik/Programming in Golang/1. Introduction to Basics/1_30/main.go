// Дано трехзначное число. Найдите сумму его цифр.
//
// Формат входных данных
// На вход дается трехзначное число.
//
// Формат выходных данных
// Выведите одно целое число - сумму цифр введенного числа.

package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	third_digit := x % 10
	second_digit := (x / 10) % 10
	first_digit := (x / 100) % 10
	result := first_digit + second_digit + third_digit
	fmt.Print(result)
}
