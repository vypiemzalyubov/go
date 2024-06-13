// Из данного числа выбросите цифры 5 и 7, при этом порядок остальных цифр не меняется.
//
// Входные данные
// Вводится одно натуральное число n, не превышающее 10000.
//
// Выходные данные
// Вывести число n, без цифр 5 и 7.
//
// Примечание: гарантируется, что число содержит хотя бы 1 символ, отличный от 5 и 7.

package main

import (
	"fmt"
)

func main() {
	n, result := 0, ""
	_, _ = fmt.Scan(&n)

	for n > 0 {
		lastD := n % 10
		if lastD != 5 && lastD != 7 {
			result = fmt.Sprintf("%d%s", lastD, result)
		}
		n /= 10
	}

	fmt.Print(result)
}
