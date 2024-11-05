// Вам необходимо написать функцию calculator следующего вида:
// func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
//
// Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.
// - в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный) канал вы должны отправить квадрат аргумента.
// - в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный) канал вы должны отправить результат умножения аргумента на 3.
// - в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
//
// Функция calculator должна быть неблокирующей, сразу возвращая управление.
// Ваша функция получит всего одно значение в один из каналов - получили значение, обработали его, завершили работу.
//
// После завершения работы необходимо освободить ресурсы, закрыв выходной канал, если вы этого не сделаете, то превысите предельное время работы.

package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	outputChan := make(chan int)

	go func() {
		defer close(outputChan)

		select {
		case num := <-firstChan:
			outputChan <- num * num
		case num := <-secondChan:
			outputChan <- num * 3
		case <-stopChan:
		}
	}()

	return outputChan
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	go func() {
		firstChan <- 4
	}()

	go func() {
		secondChan <- 5
	}()

	resultChan := calculator(firstChan, secondChan, stopChan)

	for result := range resultChan {
		fmt.Println(result)
	}
}
