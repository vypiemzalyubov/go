package task1

import (
	"sort"
)

func IntersectionNLogN(a, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)

	result := []int{}
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			result = append(result, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

func IntersectionN(a, b []int) []int {
	tmpMap := make(map[int]int)
	result := []int{}

	for _, num := range a {
		tmpMap[num]++
	}

	for _, num := range b {
		if tmpMap[num] > 0 {
			result = append(result, num)
			tmpMap[num]--
		}
	}

	return result
}
