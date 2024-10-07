// Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch {
	case n < 10 && n >= 0:
		fmt.Print(n)
	case n >= 10 && n < 100:
		fmt.Print(n / 10)
	case n >= 100 && n < 1000:
		fmt.Print(n / 100)
	case n >= 1000 && n < 10000:
		fmt.Print(n / 1000)
	case n == 10000:
		fmt.Print(n / 10000)
	}
}
