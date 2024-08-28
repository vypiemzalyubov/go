package task

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		a, b     []int
		expected []int
	}{
		{a: []int{}, b: []int{}, expected: []int{}},                         // Пустые слайсы
		{a: []int{1, 2, 3}, b: []int{4, 5, 6}, expected: []int{}},           // Нет пересечений
		{a: []int{1, 2, 3, 4}, b: []int{3, 4, 5, 6}, expected: []int{3, 4}}, // Частичное пересечение
		{a: []int{1, 2, 3}, b: []int{1, 2, 3}, expected: []int{1, 2, 3}},    // Полное пересечение
		{a: []int{1, 1, 2, 2}, b: []int{2, 2, 3, 3}, expected: []int{2}},    // Пересечение с дубликатами
		{a: []int{1, 2, 3}, b: []int{3, 2, 1}, expected: []int{1, 2, 3}},    // Пересечение с разным порядком
	}

	for _, tt := range tests {
		actual := Intersection(tt.a, tt.b)
		sort.Ints(actual)      // Сортируем фактический результат
		sort.Ints(tt.expected) // Сортируем ожидаемый результат
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Intersection(%v, %v) = %v; expected %v", tt.a, tt.b, actual, tt.expected)
		}
	}
}
