// Требуется определить, бьет ли конь, стоящий на клетке с указанными координатами (номер строки и номер столбца), фигуру, стоящую на другой указанной клетке.
//
// Формат входных данных
// Вводятся четыре числа: координаты коня и координаты другой фигуры. Все координаты — целые числа в интервале от 1 до 8.
//
// Формат выходных данных
// Программа должна вывести "YES", если конь может побить фигуру за 1 ход, в противном случае вывести "NO".

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)

	if math.Abs(x1-x2)*math.Abs(y1-y2) == 2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}