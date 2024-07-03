// На стандартный вход подается два целых числа, каждое с новой строки. Вам нужно их считать и вывести их сумму.

package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Print(x + y)
}
