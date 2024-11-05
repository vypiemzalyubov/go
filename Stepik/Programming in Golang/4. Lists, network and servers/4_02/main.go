// Вам нужно реализовать функцию, которая принимает list и переворачивает порядок его элементов,
// так чтобы последний элемент стал первым, предпоследний — вторым, и так далее.
//
// Писать функцию main не нужно!

package main

import (
	"container/list"
)

func ReverseList(l *list.List) *list.List {
	reversedList := list.New()

	for temp := l.Front(); temp != nil; temp = temp.Next() {
		reversedList.PushFront(temp.Value)
	}

	return reversedList
}
