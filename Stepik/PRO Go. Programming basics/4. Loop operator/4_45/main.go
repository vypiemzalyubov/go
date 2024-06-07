// Найдите разницу между максимальным и минимальным числами.
//
// Входные данные
// Вводится натуральное число N(N≤10**5), а затем N чисел, каждое из которых по модулю не превышает 10**6.
//
// Выходные данные
// Выведите разницу между максимальным и минимальным числами.

package main

import (
	"fmt"
)

func main() {
	var n, tmp int
	_, _ = fmt.Scan(&n, &tmp)
	maxD, minD := tmp, tmp

	for i := 0; i <= n; i++ {
		_, _ = fmt.Scan(&tmp)

		if tmp > maxD {
			maxD = tmp
		} else if tmp < minD {
			minD = tmp
		}
	}

	fmt.Print(maxD - minD)
}
