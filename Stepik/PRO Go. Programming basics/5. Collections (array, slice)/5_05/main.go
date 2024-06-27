// Дан массив, состоящий из целых чисел. Напишите программу, которая определяет, есть ли в массиве пара соседних элементов с одинаковыми знаками.
//
// Входные данные
// Сначала задано число N — количество элементов в массиве (1≤N≤10000). Далее через пробел записаны N чисел — элементы массива.
// Массив состоит из целых ненулевых чисел, каждое из которых по модулю не превышает 10000.
//
// Выходные данные
// Необходимо вывести "YES", если существует пара соседних элементов с одинаковыми знаками. В противном случае следует вывести "NO".

package main

import (
	"fmt"
)

func main() {
	n, flag := 0, false
	_, _ = fmt.Scan(&n)
	array := make([]int, n)
	_, _ = fmt.Scan(&array[0])

	for i := 1; i <= n-1; i++ {
		_, _ = fmt.Scan(&array[i])

		if array[i] > 0 && array[i-1] > 0 || array[i] < 0 && array[i-1] < 0 {
			flag = true
			break
		}
	}

	if flag {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}