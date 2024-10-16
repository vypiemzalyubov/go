// Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A.
// Если А не является числом Фибоначчи, выведите число -1.
//
// Входные данные
// Вводится натуральное число.
//
// Выходные данные
// Выведите ответ на задачу.

package main

import (
	"fmt"
)

func main() {
	var A int
	fmt.Scan(&A)

	if A < 1 {
		fmt.Println(-1)
		return
	}

	fib1, fib2 := 1, 1
	n := 2

	if A == fib1 {
		fmt.Println(1)
		return
	} else if A == fib2 {
		fmt.Println(2)
		return
	}

	for {
		nextFib := fib1 + fib2
		n++

		if nextFib == A {
			fmt.Println(n)
			return
		} else if nextFib > A {
			fmt.Println(-1)
			return
		}

		fib1, fib2 = fib2, nextFib
	}
}
