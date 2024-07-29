// Дано два натуральных числа, не заканчивающиеся на 0. Замените каждое число на обратное и вычислите их сумму.
// Например, дается два числа 624 и 1024. Каждое заменяем на обратное, то есть 624⇒426,1024⇒4201.
// Затем находим их сумму: 426+4201=4627.
//
// Входные данные
// Вводятся два натуральных числа, каждый из которых не превосходит 10**8. Гарантируется, что каждое число не оканчивается на 0.
//
// Выходные данные
// Выведите сумму обратных чисел.
//
// Примечание: предполагается задачу решить с помощью функции.

package main

import (
	"fmt"
)

func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)

	fmt.Println(reverseDigit(n) + reverseDigit(m))
}

func reverseDigit(number int) int {
	digit := 0
	for number > 0 {
		digit = digit*10 + number%10
		number /= 10
	}
	return digit
}