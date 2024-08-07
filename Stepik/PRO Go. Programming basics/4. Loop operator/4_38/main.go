// Найдите все числа от 1 до n кратные c, но НЕ кратные d.
//
// Входные данные
// Вводится 3 натуральных числа n,c,d, каждое из которых не превышает 10000. Каждое число вводится на отдельной строке.
//
// Выходные данные
// Вывести все числа от 1 до n включительно, кратные c, но НЕ кратные d.

package main

import (
	"fmt"
)

func main() {
	n, c, d := 0, 0, 0
	fmt.Scan(&n, &c, &d)

	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
		}
	}
}
