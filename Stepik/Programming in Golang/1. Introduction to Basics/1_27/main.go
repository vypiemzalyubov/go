// На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.
// Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.

package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	maxD := array[0]
	for _, v := range array {
		if v > maxD {
			maxD = v
		}
	}

	fmt.Println(maxD)
}
