// Дан массив, состоящий из целых чисел. Напишите программу, которая меняет местами первый минимальный и последний максимальный элементы массива. Индексация начинается с нуля.
//
// Входные данные
// Сначала задано число N — количество элементов в массиве (1≤N≤1000). Далее через пробел записаны N чисел — элементы массива.
// Массив состоит из целых чисел, каждое из которых не превышает 10000.
//
// Выходные данные
// Необходимо вывести массив, в котором поменяны местами первый минимальный и последний максимальный элементы массива.

package main

import (
	"fmt"
)

func main() {
	n := 0
	_, _ = fmt.Scan(&n)
	array := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&array[i])
	}

	indexMin, indexMax := 0, 0
	minD, maxD := array[0], array[0]
	for i := 1; i < n; i++ {
		if array[i] < minD {
			minD = array[i]
			indexMin = i
		} else if array[i] >= maxD {
			maxD = array[i]
			indexMax = i
		}
	}

	for index, value := range array {
		if index == indexMin {
			fmt.Print(maxD, " ")
		} else if index == indexMax {
			fmt.Print(minD, " ")
		} else {
			fmt.Print(value, " ")
		}
	}
}
