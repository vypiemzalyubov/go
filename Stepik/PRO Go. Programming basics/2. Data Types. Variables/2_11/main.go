// Напишите программу, которая считывает три целых числа и выводит на экран их произведение. Каждое число записано в отдельной строке.
//
// Формат входных данных
// На вход программе подаётся три целых числа, каждое на отдельной строке.
//
// Формат выходных данных
// Программа должна вывести произведение введенных чисел.

package main

import (
	"fmt"
)

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Print(x * y * z)
}
