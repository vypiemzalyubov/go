// Найдите самое большее число на отрезке от a до b, кратное 7.
//
// Входные данные
// Вводится два целых числа a и b(a≤b) , каждое из которых по модулю не превышает 2∗10**4.
//
// Выходные данные
// Найдите самое большее число на отрезке от a до b, кратное 7, или выведите "NO" - если таковых нет.
//
// Примечание: 0 кратен любому числу!

package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	maxD := a

	for a <= b {
		if a%7 == 0 && a > maxD {
			maxD = a
		}
		a++
	}

	if maxD%7 == 0 {
		fmt.Print(maxD)
	} else {
		fmt.Print("NO")
	}
}
