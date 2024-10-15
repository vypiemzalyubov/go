// Найдите количество минимальных элементов в последовательности.
//
// Входные данные
// Вводится натуральное число N, а затем N целых чисел последовательности.
//
// Выходные данные
// Выведите количество минимальных элементов последовательности.

package main

import "fmt"

func main() {
	var n, x, count, minD int
	fmt.Scan(&n, &x)
	minD = x
	arr := []int{x}

	for i := 1; i < n; i++ {
		fmt.Scan(&x)
		arr = append(arr, x)

		if x < minD {
			minD = x
		}
	}

	for _, v := range arr {
		if v == minD {
			count++
		}
	}

	fmt.Println(count)
}
