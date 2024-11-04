// Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.
// Функция должна называться task2().

package main

import "fmt"

func main() {
	c := make(chan string)
	go task2(c, "kek")
}

func task2(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("%v ", s)
	}
}
