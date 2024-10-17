// По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".
//
// Входные данные
// Дано число n (0<n<100).
//
// Выходные данные
// Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov.
// Между числом и словом должен стоять ровно один пробел.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch {
	case n%10 == 1 && n != 11:
		fmt.Printf("%d korova", n)
	case n%10 == 2 && n != 12 || n%10 == 3 && n != 13 || n%10 == 4 && n != 14:
		fmt.Printf("%d korovy", n)
	default:
		fmt.Printf("%d korov", n)
	}
}