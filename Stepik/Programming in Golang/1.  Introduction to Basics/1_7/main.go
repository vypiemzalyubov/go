// 1.7 Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).
//     На вход дается натуральное число, не превосходящее 10000.
//     Выведите одно целое число - число десятков.

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Print(a % 100 / 10)
}
