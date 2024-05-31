// Найдите первое число от 1 до n, кратное c, но НЕ кратное d.
//
// Входные данные
// Вводится 3 натуральных числа n,c,d, каждое из которых не превышает 10000
//
// Выходные данные
// Вывести первое число от 1 до n кратное c, но НЕ кратное d.
//
// Примечание: если такого числа нет, то ничего не надо выводить!

package main

import (
	"fmt"
)

func main() {
	n, c, d := 0, 0, 0
	_, _ = fmt.Scan(&n, &c, &d)

	for i := 1; i < n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Print(i)
			break
		}
	}
}
