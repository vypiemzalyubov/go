// Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль. Ввод данных уже написан за вас.

package main

import "fmt"

func main() {
	var x1, x2 int = 2, 2
	test(&x1, &x2)
}

func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}