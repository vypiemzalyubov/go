// По данным натуральным n и k вычислите значение Сnk = n!/k!(n−k)! (число сочетаний из n элементов по k).
//
// Входные данные
// На первой строке вводится натуральное число n, не превосходящее 10.
// На второй строке вводится целое неотрицательное число k, не превосходящее 10.
//
// Выходные данные
// Необходимо вывести значение Сnk.
//
// Примечание: данную задачу предполагается решить с помощью функций. Не забывайте, что 0!=1.

package main

import "fmt"

func main() {
	var n, k int
	_, _ = fmt.Scan(&n, &k)

	factN := factorial(n)
	factK := factorial(k)
	factNSubtractK := factorial(n - k)

	fmt.Println(factN / (factK * factNSubtractK))
}

func factorial(number int) int {
	fact := 1
	for i := 1; i <= number; i++ {
		fact = fact * i
	}
	return fact
}
