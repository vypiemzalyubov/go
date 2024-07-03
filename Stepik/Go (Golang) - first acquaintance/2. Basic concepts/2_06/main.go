// На стандартный вход подается три целых числа, в одной строке через пробел. Вам нужно их считать и вывести их сумму.

package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Print(x + y + z)
}
