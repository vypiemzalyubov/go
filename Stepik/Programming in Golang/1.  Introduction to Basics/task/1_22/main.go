// Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается.
// Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.
// Программа получает на вход три натуральных числа: x, p, y.
// Программа должна вывести одно целое число.

package main

import "fmt"

func main() {
	var x, p, y, year int
	fmt.Scan(&x, &p, &y)
	for x <= y {
		x = x + x*p/100
		year++
	}
	fmt.Print(year)
}
