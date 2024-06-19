// Дан массив, состоящий из целых чисел. Напишите программу, которая определяет является ли массив палиндромом.
// То есть если перевернуть массив, то получится массив, равный первоначальному.
//
// Входные данные
// Сначала задано число N — количество элементов в массиве (1≤N≤1000). Далее через пробел записаны N чисел — элементы массива.
// Массив состоит из целых чисел, каждое из которых не превышает 100000.
//
// Выходные данные
// Необходимо вывести "YES", если массив является палиндромом, в противном случае - "NO".

package main

import (
	"fmt"
)

func main() {
	n, flag := 0, false
	_, _ = fmt.Scan(&n)
	array := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&array[i])
	}

	for i := 0; i < n; i++ {
		if array[i] == array[n-1-i] {
			flag = true
		} else {
			flag = false
			break
		}
	}

	if flag {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
