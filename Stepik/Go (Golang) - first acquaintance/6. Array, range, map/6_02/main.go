// Есть массив nums. Вам нужно написать цикл для перебора его значений с использованием range.
// В цикле нужно посчитать сумму всех элементов массива. Для этого объявлена переменная sum.

sum := 0
for _, v := range nums {
	sum += v
}