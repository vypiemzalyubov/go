// Даны два предложения. Найти общее количество букв 'b' в них.
//
// Входные данные
// Вводятся две строки, каждая из которых по длине не превосходит 1000.
//
// Выходные данные
// Выведите  общее количество букв 'b' в них.
//
// Примечание: предполагается задачу решить с помощью функции.

package main

import (
	"fmt"
)

func main() {
	var word1, word2 string
	_, _ = fmt.Scan(&word1, &word2)

	fmt.Println(bCount(word1) + bCount(word2))
}

func bCount(word string) int {
	cnt := 0
	for i := 0; i < len(word); i++ {
		if word[i] == 'b' {
			cnt++
		}
	}

	return cnt
}
