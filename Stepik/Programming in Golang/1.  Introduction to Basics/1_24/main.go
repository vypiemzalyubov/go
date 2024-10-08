// На вход подается число типа float64. Вам нужно вывести преобразованное число по правилу: число возводится в квадрат,
// затем обрезается дробная часть так что остается 4 знака после запятой.
// Но если число равно или меньше нуля - выводить:
// "число R не подходит", где R - исходное число с 2 цифрами после запятой и с 2 по ширине.
// А если число больше чем 10 000 - выводить исходное число в экспоненциальном представлении (знак экспоненты в нижнем регистре).

package main

import "fmt"

func main() {
	var n, result float64
	fmt.Scan(&n)
	result = n * n
	if n < 0 || n == 0 {
		fmt.Printf("число %.2f не подходит", n)
	} else if n > 10000 {
		fmt.Printf("%e", n)
	} else {
		fmt.Printf("%.4f", result)
	}
}
