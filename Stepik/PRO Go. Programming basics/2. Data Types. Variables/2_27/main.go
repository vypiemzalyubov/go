// Даны два числа. Найти их среднее арифметическое.
//
// Формат входных данных
// На вход дается два целых положительных числа a и b, каждое из которых не превышает 10000.
//
// Формат выходных данных
// Программа должна вывести среднее арифметическое чисел a и b.

package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Print((a + b) / 2)
}
