package task1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		a        []int
		b        []int
		expected []int
	}{
		{a: []int{}, b: []int{}, expected: []int{}},
		{a: []int{100, 101}, b: []int{}, expected: []int{}},
		{a: []int{}, b: []int{-8, 1000, 33}, expected: []int{}},
		{a: []int{-1, 2, -3}, b: []int{-10, 0, -15}, expected: []int{}},
		{a: []int{1, 2, 3, 4}, b: []int{3, 4, 5, 6}, expected: []int{3, 4}},
		{a: []int{91, 65, 38, 0, 13}, b: []int{91, 65, 38, 0, 13}, expected: []int{91, 65, 38, 0, 13}},
		{a: []int{10, 11, 12}, b: []int{12, 11, 10}, expected: []int{10, 11, 12}},
		{a: []int{1, 1, 2, 2}, b: []int{1, 2, 2}, expected: []int{1, 2, 2}},
	}

	for _, tt := range tests {
		actualNLogN := IntersectionNLogN(tt.a, tt.b)
		assert.ElementsMatch(t, tt.expected, actualNLogN, "IntersectionNLogN(%v, %v)", tt.a, tt.b)

		actualN := IntersectionN(tt.a, tt.b)
		assert.ElementsMatch(t, tt.expected, actualN, "IntersectionN(%v, %v)", tt.a, tt.b)
	}
}
