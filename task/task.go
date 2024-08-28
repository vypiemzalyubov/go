package task

func Intersection(a, b []int) []int {
	// Используем map для хранения уникальных элементов из первого слайса
	elements := make(map[int]struct{})
	result := []int{}

	// Добавляем элементы из первого слайса в map
	for _, num := range a {
		elements[num] = struct{}{}
	}

	// Проверяем, какие элементы из второго слайса есть в map
	for _, num := range b {
		if _, found := elements[num]; found {
			result = append(result, num)
			delete(elements, num) // Удаляем элемент из map, чтобы избежать дубликатов в результате
		}
	}

	return result
}
