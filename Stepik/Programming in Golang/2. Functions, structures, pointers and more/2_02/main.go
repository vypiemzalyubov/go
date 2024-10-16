// Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
//
// Входные данные
// Вводится четыре числа.
//
// Выходные данные
// Необходимо вернуть из функции наименьшее из 4-х данных чисел

package main

import (
	"fmt"
	"slices"
)

func main() {
	minimumFromFour()
}

func minimumFromFour() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	s := []int{a, b, c, d}

	return slices.Min(s)
}
