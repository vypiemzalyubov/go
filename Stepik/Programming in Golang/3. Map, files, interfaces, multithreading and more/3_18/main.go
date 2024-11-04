// Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
// Функция должна называться task().

package main

func main() {
	c := make(chan int)
	go task(c, 5)
}

func task(c chan int, n int) {
	c <- n + 1
}
