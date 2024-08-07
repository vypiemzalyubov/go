// Напишите программу, которая в последовательности трехзначных чисел находит количество всех чисел сумма цифр которых равна 8.
// В ответе запишите найденное количество чисел.

package main

import (
	"fmt"
)

func main() {
	cnt := 0
	for i := 100; i <= 999; i++ {
		sum := 0
		for k := i; k > 0; k /= 10 {
			sum += k % 10
		}
		if sum == 8 {
			cnt++
		}
	}
	fmt.Print(cnt)
}
