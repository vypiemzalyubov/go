// Дано натуральное число N. Выведите его представление в двоичном виде.
//
// Входные данные
// Задано единственное число N
//
// Выходные данные
// Необходимо вывести требуемое представление числа N.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, 0, n)

	for n > 0 {
		arr = append(arr, n%2)
		n /= 2
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i])
	}
}
