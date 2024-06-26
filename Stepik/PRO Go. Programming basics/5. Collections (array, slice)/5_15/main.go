// Проверьте, является ли двумерный массив симметричным относительно главной диагонали.
// Главная диагональ — та, которая идёт из левого верхнего угла двумерного массива в правый нижний.
//
// Входные данные
// Программа получает на вход число n, n≤100, являющееся числом строк и столбцов в массиве.
// Далее во входном потоке идет n строк по n чисел разделенных пробелами, являющихся элементами массива.
//
// Выходные данные
// Программа должна выводить "YES" для симметричного массива и "NO" для несимметричного.

package main

import (
	"fmt"
)

func main() {
	n, flag := 0, "YES"
	_, _ = fmt.Scan(&n)

	array := make([][]int, n)
	for i := range array {
		array[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Scan(&array[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if array[i][j] != array[j][i] {
				flag = "NO"
				break
			}
		}
	}

	fmt.Print(flag)
}
