// Определите, принадлежит ли точка x  одному из выделенных отрезков B или D.
//
// Входные данные
// Дано одно целое число x, по модулю не превышающее 10000.
//
// Выходные данные
// Выведите "YES", если точка принадлежит одному из выделенных отрезков B или D (включая границы), в противном случае - "NO".

package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	if x > -4 && x < 2 || x > 4 && x < 10 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
