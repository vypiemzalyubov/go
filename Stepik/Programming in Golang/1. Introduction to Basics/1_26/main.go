// Напишите программу, принимающая на вход число N(N≥4), а затем N целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := []int{}

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		s = append(s, x)
	}

	fmt.Println(s[3])
}
