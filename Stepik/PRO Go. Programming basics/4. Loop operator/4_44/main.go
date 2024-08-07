// Найдите максимальное из введенных чисел.
//
// Входные данные
// Вводится натуральное число N(N≤10**5), а затем N целых чисел, каждое из которых по модулю не превышает 10**6.
//
// Выходные данные
// Выведите максимальное из введенных чисел.

package main

import (
	"fmt"
)

func main() {
	var n, tmp int
	_, _ = fmt.Scan(&n, &tmp)
	maxD := tmp

	for i := 0; i <= n; i++ {
		_, _ = fmt.Scan(&tmp)

		if tmp > maxD {
			maxD = tmp
		}
	}

	fmt.Print(maxD)
}
