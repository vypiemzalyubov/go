// Кажется, еще совсем недавно наступил новый 2013 год. А знали ли Вы забавный факт о том,
// что 2013 год является первым годом после далекого 1987 года, в котором все цифры различны?
// Теперь же Вам предлагается решить следующую задачу: задан номер года, найдите наименьший номер года, который строго больше заданного и в котором все цифры различны.
//
// Входные данные
// В единственной строке задано целое число y (1000 ≤ y ≤ 9000) — номер года.
//
// Выходные данные
// Выведите единственное целое число — минимальный номер года, который строго больше y, в котором все цифры различны. Гарантируется, что ответ существует.

package main

import (
	"fmt"
)

func main() {
	var y int
	_, _ = fmt.Scan(&y)

	for y > 0 {
		firstD := (y + 1) / 1000 % 10
		secondD := (y + 1) / 100 % 10
		thirdD := (y + 1) / 10 % 10
		fourthD := (y + 1) % 10
		if firstD != secondD && firstD != thirdD && firstD != fourthD && secondD != thirdD && secondD != fourthD && thirdD != fourthD {
			fmt.Print(y + 1)
			break
		}
		y++
	}
}
