// 1.21 Напишите программу, которая считывает целые числа с консоли по одному числу в строке.
//      Для каждого введённого числа проверить:
//      если число меньше 10, то пропускаем это число;
//      если число больше 100, то прекращаем считывать числа;
//      в остальных случаях вывести это число обратно на консоль в отдельной строке.

package main

import "fmt"

func main() {
	var number int
	for {
		fmt.Scan(&number)
		if number < 10 {
			continue
		} else if number > 100 {
			break
		} else {
			fmt.Println(number)
		}
	}
}
