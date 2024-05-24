// Требуется определить, бьет ли ферзь, стоящий на клетке с указанными координатами (номер строки и номер столбца), фигуру, стоящую на другой указанной клетке.
//
// Формат входных данных
// Вводятся четыре числа: координаты ферзя и координаты другой фигуры. Координаты — целые числа в интервале от 1 до 8.
//
// Формат выходных данных
// Требуется вывести "YES", если ферзь может побить фигуру за 1 ход, в противном случае - "NO".

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)

	if x1 == x2 || y1 == y2 || math.Abs(x1-x2) == math.Abs(y1-y2) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
