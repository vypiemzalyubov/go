// Найдите минимальное среди последовательности целых чисел.
//
// Входные данные
// Вводится натуральное число N(N≤10**5), а затем N чисел, каждое из которых по модулю не превышает 10**6.
//
// Выходные данные
// Выведите минимальное из введенных чисел.

package main

import (
	"fmt"
)

func main() {
	n, number, cnt := 0, 0, 0
	_, _ = fmt.Scan(&n, &number)
	minD := number

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&number)

		if number < minD {
			minD = number
			cnt = 1
		} else if number == minD {
			cnt++
		}
	}

	fmt.Print(cnt)
}
