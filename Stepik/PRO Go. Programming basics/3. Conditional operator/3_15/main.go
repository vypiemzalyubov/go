// Требуется определить, бьет ли слон, стоящий на клетке с указанными координатами (номер строки и номер столбца), фигуру, стоящую на другой указанной клетке.
//
// Входные данные
// Вводятся четыре числа: координаты слона и координаты другой фигуры. Координаты — целые числа в интервале от 1 до 8.
//
// Выходные данные
// Требуется вывести слово "YES", если слон способен побить фигуру за 1 ход, в противном случае вывести слово "NO".

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64
	fmt.Scan(&x1, &x2, &y1, &y2)

	if math.Abs(y1-x1) == math.Abs(y2-x2) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
