// Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения.
//
// Функция work() ничего не принимает и не возвращает.

package main

import (
	"fmt"
	"time"
)

func work() {
	time.Sleep(2 * time.Second)
	fmt.Println("Работа завершена")
}

func main() {
	done := make(chan bool)

	go func() {
		work()
		done <- true
	}()

	<-done
	fmt.Println("Главная функция завершена")
}
