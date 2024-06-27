// Дано натуральное число N. Выведите количество единиц в его двоичном представлении.
// Например: N=11. В двоичном системе счисления это число будет выглядеть следующим образом: 1011. Количество единиц в нем будет равно 3.
//
// Входные данные
// Вводится натуральное число N, не превышающее 10**9.
//
// Выходные данные
// Выведите количество единиц в двоичном представлении числа N.

package main

import (
	"fmt"
)

func main() {
	n, cnt := 0, 0
	fmt.Scan(&n)

	for ; n > 0; n /= 2 {
		cnt += n % 2
	}

	fmt.Print(cnt)
}