// Дано два натуральных числа. Найдите количество разрядов каждого из них и выведите их произведение.
//
// Входные данные
// Вводятся два натуральных числа - n, m, каждое из которых не превосходит 10**9.
//
// Выходные данные
// Выведите произведение количества разрядов данных чисел.
//
// Примечание: предполагается задачу решить с помощью функции.

package main

import (
	"fmt"
)

func main() {
	var n, m string
	_, _ = fmt.Scan(&n, &m)

	fmt.Print(digitsNumber(n) * digitsNumber(m))
}

func digitsNumber(number string) int {
	return int(len(number))
}
