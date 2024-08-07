// Напишите программу, которая определяет, является ли заданное натуральное число трёхзначным.
//
// Формат входных данных
// На вход подается одно натуральное число.
//
// Формат выходных данных
// Выведите "YES", если все введенное число является трёхзначным, в противном случае - "NO".

package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	if x > 99 && x < 1000 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
