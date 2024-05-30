// Посчитать, сколько раз встречается цифра 7 в последовательности чисел от 1 до N, включая N.
//
// Входные данные
// В первой строке находится единственное натуральное число N, не превышающее 10**9.
//
// Выходные данные
// Выведите  сколько раз встречается цифра 7 в последовательности чисел от 1 до N.

package main

import (
	"fmt"
)

func main() {
	n, cnt := 0, 0
	_, _ = fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		k := i
		for ; k > 0; k /= 10 {
			if k%10 == 7 {
				cnt++
			}
		}
	}
	fmt.Print(cnt)
}
