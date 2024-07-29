// Даны две последовательности. Первая: 1,2,3,...,n, а вторая: 1,2,3,...,m. Найдите средние арифметические обеих последовательностей и выведите их сумму.
//
// Входные данные
// Вводятся два натуральных числа - n, m, каждое из которых не превосходит 1000.
//
// Выходные данные
// Выведите сумму двух средних арифметических.
//
// Примечание: предполагается задачу решить с помощью функции.

package main

import (
	"fmt"
)

func main() {
	var n, m float64
	_, _ = fmt.Scan(&n, &m)

	fmt.Println(arithmeticAvg(n) + arithmeticAvg(m))
}

func arithmeticAvg(number float64) float64 {
	var sum int
	for i := 1; i <= int(number); i++ {
		sum += i
	}

	return float64(sum) / number
}
