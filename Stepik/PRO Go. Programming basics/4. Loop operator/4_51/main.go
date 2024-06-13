// Для целых чисел a и n вычислить a**n.
//
// Входные данные
// Вводятся два целых числа a и n , при этом 1≤a<100,0≤n≤20.
//
// Выходные данные
// Выводится одно число a**n.
//
// Примечание: гарантируется, что ответ не больше 2⋅10**9. НЕ используйте Math.Pow!

package main

import (
	"fmt"
)

func main() {
	var a, n int
	_, _ = fmt.Scan(&a, &n)
	amount := 1

	for i := 0; i < n; i++ {
		amount *= a
	}

	fmt.Print(amount)
}
