// Дано натуральное число, выведите его последнюю цифру.
// На вход дается натуральное число N, не превосходящее 10000.
// Выведите одно целое число - ответ на задачу.

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	last_digit := a % 10
	fmt.Print(last_digit)
}
