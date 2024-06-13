// Найдите количество минимальных элементов в последовательности.
//
// Входные данные
// Вводится натуральное число N(N≤10**5), а затем N чисел. Все числа по модулю не превышают 10**6.
//
// Выходные данные
// Выведите количество минимальных элементов.

package main

import (
	"fmt"
)

func main() {
	var n, number int
	_, _ = fmt.Scan(&n, &number)
	minD := number
	cnt := 1

	for i := 1; i < n; i++ {
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
