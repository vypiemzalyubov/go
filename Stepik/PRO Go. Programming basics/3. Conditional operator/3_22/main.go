// По данному месяцу, определите количество дней в этом месяце.
//
// Входные данные
// Вводится одно число - номера месяца, принимающее значения от 1 до 12.
//
// Выходные данные
// Выведите одно число - количество дней в этом месяце.

package main

import (
	"fmt"
)

func main() {
	var month int
	fmt.Scan(&month)

	switch month {
	case 1:
		fmt.Print(31)
	case 2:
		fmt.Print(29)
	case 3:
		fmt.Print(31)
	case 4:
		fmt.Print(30)
	case 5:
		fmt.Print(31)
	case 6:
		fmt.Print(30)
	case 7:
		fmt.Print(31)
	case 8:
		fmt.Print(31)
	case 9:
		fmt.Print(30)
	case 10:
		fmt.Print(31)
	case 11:
		fmt.Print(30)
	case 12:
		fmt.Print(31)
	}
}
