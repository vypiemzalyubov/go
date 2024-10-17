// Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).
//
// Входные данные
// Сначала задано число N — количество элементов в массиве (1≤N≤100). Далее через пробел записаны N чисел — элементы массива. Массив состоит из целых чисел.
//
// Выходные данные
// Необходимо вывести все элементы массива с чётными индексами.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	array := []int{}

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		array = append(array, x)
		if i%2 == 0 {
			fmt.Printf("%v ", array[i])
		}
	}
}