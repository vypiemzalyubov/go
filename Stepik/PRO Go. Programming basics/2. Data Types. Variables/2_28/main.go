// Дано положительное вещественное число x. Выведите его дробную часть.
//
// Входные данные
// Вводится положительное вещественное число.
//
// Выходные данные
// Дробная часть входного числа.

package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Scan(&x)
	fmt.Print(x - float64(int(x)))
}
