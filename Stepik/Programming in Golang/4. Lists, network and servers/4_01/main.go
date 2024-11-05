// Реализуйте функции для очереди FIFO (First In, First Out (ПЕРВЫЙ пришел, ПЕРВЫЙ вышел)) с помощью списков.
// Должны быть данные функции:
// - Push (добавление элемента)
// - Pop (удаление элемента и его возврат)
// - printQueue (печать очереди в одну строку без пробелов)
//
// Функцию main писать не нужно! Писать код вне функций не нужно.

package main

import (
	"container/list"
	"fmt"
)

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	elem := queue.Remove(queue.Front())
	return elem
}

func printQueue(queue *list.List) {
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Printf("%v", temp.Value)
	}
}
