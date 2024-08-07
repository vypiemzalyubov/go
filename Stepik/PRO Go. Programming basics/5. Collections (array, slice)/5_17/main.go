// Даны два числа n и m. Создайте двумерный массив A[n,m] и заполните его по следующим правилам: Числа, стоящие в строке 0 или в столбце 0 равны 1.
// Для всех остальных элементов массива- элемент равен сумме двух элементов, стоящих слева и сверху от него.
//
// Входные данные
// Программа получает на вход строку, в котором через пробел записаны два натуральных числа n и m через пробел, каждое из которых не превышает 10.
//
// Выходные данные
// Выведите полученный массив.

package main

import (
	"fmt"
)

func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)

	mtrx := make([][]int, n)
	for i := range mtrx {
		mtrx[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				mtrx[i][j] = 1
			} else {
				mtrx[i][j] = mtrx[i-1][j] + mtrx[i][j-1]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(mtrx[i][j], " ")
		}
		fmt.Println()
	}
}
