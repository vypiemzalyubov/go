// Совершенное число — натуральное число, равное сумме всех своих собственных делителей (то есть всех положительных делителей, отличных от самого числа).
// Например, 6 — это совершенное число, так как сумма его собственных делитей 1+2+3 равняется 6.
// Напишите программу, которая будет искать совершенные числа.
// В ответе укажите первые три совершенных числа через запятую без пробелов.

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 1000; i++ {
		sum := 0
		for k := 1; k < i; k++ {
			if i%k == 0 {
				sum += k
			}
		}
		if i == sum {
			fmt.Print(i, ",")
		}
	}
}
