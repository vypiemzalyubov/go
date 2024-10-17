// Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму.
// Пакет "fmt" уже импортирован, функция и пакет main объявлены.
//
// Пример вызова вашей функции:
// a, b := sumInt(1, 0)
// fmt.Println(a, b)
//
// Результат: 2, 1

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sumInt(n)
}

func sumInt(n ...int) (int, int) {
	var sum int

	for _, v := range n {
		sum += v
	}

	return len(n), sum
}
