// Напишите программу, которая циклически сдвигает элементы массива вправо. Например, если элементы нумеруются, начиная с нуля,
// то 0-й элемент становится 1-м, 1-й становится 2-м, ..., последний становится 0-м, то есть массив [3,5,7,9] превращается в массив [9,3,5,7]).
//
// Входные данные
// Сначала задано число N — количество элементов в массиве (1≤N≤35). Далее через пробел записаны N чисел — элементы массива.
// Массив состоит из целых чисел, каждое из которых не превышает 10000.
//
// Выходные данные
// Необходимо вывести массив, полученный после сдвига элементов.

package main

import (
	"fmt"
)

func main() {
	n := 0
	_, _ = fmt.Scan(&n)
	array := make([]int, n)

	for i := 0; i < len(array); i++ {
		_, _ = fmt.Scan(&array[i])
	}

	for i := len(array) - 1; i > 0; i-- {
		array[i], array[i-1] = array[i-1], array[i]
	}

	for i := 0; i < len(array); i++ {
		fmt.Printf("%d%s", array[i], " ")
	}
}
