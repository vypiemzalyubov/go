// Найдите сумму введенных чисел, которые кратны 2 , но не кратны 3 .
//
// Входные данные
// Вводится натуральное число N(N≤10**5), а затем N чисел, каждое из которых по модулю не превышает 2∗10**4.
//
// Выходные данные
// Вычислите сумму введенных чисел, кратных 2 , но не кратных 3.

package main

import (
	"fmt"
)

func main() {
	var n, x int
	_, _ = fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&x)

		if x%2 == 0 && x%3 != 0 {
			sum += x
		}
	}
	fmt.Print(sum)
}
