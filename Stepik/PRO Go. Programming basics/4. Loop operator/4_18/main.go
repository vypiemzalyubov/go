// Наибольший общий делитель (НОД) двух или нескольких чисел – это наибольшее число, на которое делятся заданные числа. Напишите программу, которая вычисляет НОД двух чисел.
// Для нахождения НОД-а существует алгоритм Евклида.
//
// Входные данные
// Входная строка содержит два натуральных числа, каждое из которых не превышает 10**8.
//
// Выходные данные
// Программа должна вывести одно натуральное число - НОД заданных чисел.

package main

import (
	"fmt"
)

func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)

	for a != 0 && b != 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}

	fmt.Print(a + b)
}
