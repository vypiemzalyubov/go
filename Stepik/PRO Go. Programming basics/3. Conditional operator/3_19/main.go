// По координатам точки определите принадлежности к одной из координатных четвертей.
//
// Формат входных данных
// Вводится два вещественных числа x и y - координаты точки, каждое из которых по модулю не превышает 10**7.
//
// Формат выходных данных
// Выведите одно целое число — обозначающее соответствующую четверть, к которой относится.

package main

import (
	"fmt"
)

func main() {
	var x, y float64
	fmt.Scan(&x, &y)

	if x > 0 {
		if y > 0 {
			fmt.Print(1)
		} else {
			fmt.Print(4)
		}
	} else if y > 0 {
		fmt.Print(2)
	} else {
		fmt.Print(3)
	}
}
