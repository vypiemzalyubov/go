// На вход подаются три целых числа. Необходимо сосчитать и вывести их сумму и произведение на разных строках.

package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(x + y + z)
	fmt.Println(x * y * z)
}
