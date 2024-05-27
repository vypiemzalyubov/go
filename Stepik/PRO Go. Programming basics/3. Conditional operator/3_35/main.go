// Футбольная команда набирает мальчиков от 12 до 18 лет включительно. Напишите программу, которая запрашивает возраст и пол претендента,
// используя обозначение пола буквы m (от male – мужчина) и f (от female – женщина) и определяет, подходит ли претендент для вступления в команду или нет.
// Если претендент подходит, то выведите «YES», иначе выведите «NO».
//
// Формат входных данных
// На вход программе подаётся натуральное число – возраст претендента и буква обозначающая пол m (мужчина) или f (женщина).
//
// Формат выходных данных
// Программа должна вывести текст в соответствии с условием задачи.

package main

import (
	"fmt"
)

func main() {
	var age int
	var sex string
	fmt.Scan(&age, &sex)

	switch {
	case sex == "m" && age > 11 && age < 19:
		fmt.Print("YES")
	default:
		fmt.Print("NO")
	}
}