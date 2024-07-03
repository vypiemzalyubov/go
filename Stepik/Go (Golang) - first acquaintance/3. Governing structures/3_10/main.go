// На вход подается целое число. Вам необходимо вывести произведение всех чисел от 1 до данного числа (включительно).

package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	s := 1
	for i := 1; i <= x; i++ {
		s *= i
	}
	fmt.Println(s)
}
