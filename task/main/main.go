package main

import (
	"fmt"
)

func Intersection(a, b []int) []int {
	elements := make(map[int]struct{})
	result := []int{}

	for _, num := range a {
		elements[num] = struct{}{}
	}

	for _, num := range b {
		if _, found := elements[num]; found {
			result = append(result, num)
			delete(elements, num) // Удаляем элемент, чтобы избежать дубликатов
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}

	result := Intersection(a, b)
	fmt.Println("Intersection:", result)
}
