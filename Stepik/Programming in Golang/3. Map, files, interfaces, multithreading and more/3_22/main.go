// Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельных горутинах вызвать функцию work() 10 раз
// и дождаться результатов выполнения вызванных функций.
//
// Функция work() ничего не принимает и не возвращает. Пакет "sync" уже импортирован.

package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	time.Sleep(2 * time.Second)
	fmt.Println("Работа завершена")
}

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			work()
		}(wg)
	}

	wg.Wait()
	fmt.Println("Главная функция завершена")
}
