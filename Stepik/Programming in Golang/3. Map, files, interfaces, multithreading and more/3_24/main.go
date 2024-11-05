// Вам необходимо написать функцию calculator следующего вида:
// func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
//
// В качестве аргумента эта функция получает два канала только для чтения, возвращает канал только для чтения.
// Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу.
// Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.
//
// Функция calculator должна быть неблокирующей, сразу возвращая управление.
// Выходной канал должен быть закрыт после выполнения всех оговоренных условий, если вы этого не сделаете, то превысите предельное время работы.

package main

import (
	"fmt"
)

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	outputChan := make(chan int)

	go func() {
		defer close(outputChan)

		sum := 0
		for {
			select {
			case num, ok := <-arguments:
				if !ok {
					return
				}
				sum += num
			case <-done:
				outputChan <- sum
				return
			}
		}
	}()

	return outputChan
}

func main() {
	arguments := make(chan int)
	done := make(chan struct{})

	go func() {
		arguments <- 1
		arguments <- 2
		arguments <- 3
		close(arguments)
	}()

	go func() {
		done <- struct{}{}
	}()

	resultChan := calculator(arguments, done)

	for result := range resultChan {
		fmt.Println("Сумма:", result)
	}
}
