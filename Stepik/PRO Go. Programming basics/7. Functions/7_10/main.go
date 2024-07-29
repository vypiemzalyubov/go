// Даны два натуральных числа. Выяснить, в каком из них сумма цифр больше.
//
// Входные данные
// Вводятся два натуральных числа, каждый из которых не превосходит 10**9.
//
// Выходные данные
// Выведите цифру 1, если сумма цифр первого числа больше, чем сумма цифр второго числа.
// Выведите цифру 2, если сумма цифр второго числа больше, чем сумма цифр первого числа.
// Выведите цифру 0, если сумма цифр первого числа равно сумме цифр второго числа.
//
// Примечание: предполагается задачу решить с помощью функции.

package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)

	fmt.Print(compareSum(a, b))
}

func digitsSum(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}

	return sum
}

func compareSum(number1, number2 int) int {
	switch {
	case digitsSum(number1) > digitsSum(number2):
		return 1
	case digitsSum(number1) < digitsSum(number2):
		return 2
	default:
		return 0
	}
}
