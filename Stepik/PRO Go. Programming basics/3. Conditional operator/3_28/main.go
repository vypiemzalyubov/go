// Васе задали домашнее задание сделать k задач. За день Вася успевает сделать m задач. Сколько дней нужно, чтобы Вася сделал домашнее задание?
//
// Входные данные
// Программа получает на вход числа k и m, каждое из которых не превышает 1000.
//
// Выходные данные
// Выведите ответ на задачу.

package main

import (
	"fmt"
	"math"
)

func main() {
	var k, m float64
	fmt.Scan(&k, &m)
	fmt.Print(math.Ceil(k / m))
}