// Из натурального числа удалить заданную цифру.
//
// Входные данные
// Вводятся натуральное число и цифра, которую нужно удалить.
//
// Выходные данные
// Вывести число без заданных цифр.

package main

import "fmt"

func main() {
	var x, y string
	fmt.Scan(&x, &y)

	for _, v := range x {
		if string(v) != y {
			fmt.Print(string(v))
		}
	}
}
