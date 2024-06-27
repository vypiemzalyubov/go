// В метании молота состязается n спортcменов. Каждый из них сделал m бросков. Победителем считается тот спортсмен, у которого сумма результатов по всем броскам максимальна.
// Если перенумеровать спортсменов числами от 0 до n−1, а попытки каждого из них – от 0 до m−1, то на вход программа получает массив A[n,m], состоящий из неотрицательных целых чисел.
// Программа должна определить максимальную сумму чисел в одной строке и вывести на экран эту сумму и номер строки, для которой достигается эта сумма.
//
// Входные данные
// Программа получает на вход два числа n и m, являющиеся числом строк и столбцов в массиве. Далее во входном потоке идет n строк по m чисел, являющихся элементами массива.
//
// Выходные данные
// Программа должна вывести 2 числа: сумму и номер строки, для которой эта сумма достигается. Если таких строк несколько, то выводится номер наименьшей из них.
// Не забудьте, что нумерация строк (спортсменов) начинается с 0.

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
		for j := 0; j < m; j++ {
			_, _ = fmt.Scan(&mtrx[i][j])
		}
	}

	maxSum := 0
	maxNumber := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum += mtrx[i][j]
		}
		if sum > maxSum {
			maxSum = sum
			maxNumber = i
		}
	}

	fmt.Println(maxSum)
	fmt.Println(maxNumber)
}