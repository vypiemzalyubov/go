// По данному трехзначному числу, определите все ли его цифры различны.
//
// Формат входных данных
// На вход подается одно натуральное трехзначное число.
//
// Формат выходных данных
// Выведите "YES", если все цифры числа различны, в противном случае - "NO".

package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	first_d := x / 100
	second_d := x / 10 % 10
	third_d := x % 10

	if first_d != second_d && second_d != third_d && first_d != third_d {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
