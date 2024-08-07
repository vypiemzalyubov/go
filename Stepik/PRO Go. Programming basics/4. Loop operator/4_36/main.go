// По данной последовательности строк, определите какой по счету встретится строка s.
//
// Входные данные
// Вводится строка s, а далее последовательность строк. Каждая строка по длине не превышает 30 символов.
//
// Выходные данные
// Определите каким по счету встретится строка s в данной последовательности и выведите этот номер.
//
// Примечание: гарантируется, что данная строка встретится в последовательности. После нее строки НЕ нужно считывать!

package main

import (
	"fmt"
)

func main() {
	var s, x string
	fmt.Scan(&s)

	for i := 1; x != s; i++ {
		fmt.Scan(&x)
		if x == s {
			fmt.Print(i)
			break
		}
	}
}
