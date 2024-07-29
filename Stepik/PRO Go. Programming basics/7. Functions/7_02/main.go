// Найдите значение z=sign(a)+sign(b), где
// sign(x)=⎧−1,x<0
// 		   ⎨0,x=0
// 		   ⎩1,x>0
//
// Входные данные
// Вводится строка, в которой через пробел записаны два целых числа a и b.
//
// Выходные данные
// Необходимо вывести значение z.
//
// Примечание: данную задачу предполагается решить с помощью функций.

package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)

	fmt.Print(sign(a) + sign(b))
}

func sign(number int) int {
	switch {
	case number < 0:
		return -1
	case number == 0:
		return 0
	default:
		return 1
	}
}
