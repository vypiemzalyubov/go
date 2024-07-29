// Проверьте, является ли число простым.
//
// Входные данные
// Вводится одно натуральное число n, принимающие значения от 2 до 2⋅10**5.
//
// Выходные данные
// Необходимо вывести "prime", если число простое, или "composite", если число составное.
//
// Примечание: данную задачу предполагается решить с помощью функций.

package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	fmt.Print(easyN(n))
}

func easyN(number int) string {
	flag := "prime"

	for i := 2; i <= number; i++ {
		if number%i == 0 && number != i {
			flag = "composite"
		}
	}

	return flag
}
