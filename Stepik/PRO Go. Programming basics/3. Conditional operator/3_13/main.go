// Напишите программу, которая по введённому возрасту пользователя сообщает, к какой возрастной группе он относится:
// - до 13 включительно – детство;
// - от 14 до 24 включительно – молодость;
// - от 25 до 59 включительно – зрелость;
// - от 60 – старость.
//
// Формат входных данных
// На вход программе подаётся одно целое число – возраст пользователя.
//
// Формат выходных данных
// Программа должна вывести название возрастной группы.

package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Scan(&age)

	if age < 14 {
		fmt.Print("детство")
	} else if age > 14 && age < 25 {
		fmt.Print("молодость")
	} else if age > 24 && age < 60 {
		fmt.Print("зрелость")
	} else {
		fmt.Print("старость")
	}
}