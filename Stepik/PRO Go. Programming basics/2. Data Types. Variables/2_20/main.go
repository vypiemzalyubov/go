// Дано натуральное число, выведите его последнюю цифру.
//
// Формат входных данных
// На вход дается натуральное число N, не превосходящее 10000.
//
// Формат выходных данных
// Выведите последнюю цифру числа N.

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(n % 10)
}