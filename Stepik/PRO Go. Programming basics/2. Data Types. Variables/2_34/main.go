// Дано натуральное целое число. Найдите число сотен(то есть третью справа цифру).
//
// Формат входных данных
// На вход дается натуральное число, не превосходящее 10000.
//
// Формат выходных данных
// Выведите одно целое число - число сотен.

package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Print(x / 100 % 10)
}
