// Последовательность состоит из целых чисел и завершается числом 0.
// Определите, сколько раз в этой последовательности меняется знак. Гарантируется, что последовательность не пустая.
//
// Входные данные
// Вводится последовательность целых чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит). Числа по модулю не превышают 1000.
//
// Выходные данные
// Выведите сколько раз в последовательности меняется знак.

package main

import (
	"fmt"
)

func main() {
	n, cnt := 0, 0
	fmt.Scan(&n)

	for n != 0 {
		tmp := n
		fmt.Scan(&n)
		if tmp > 0 && n < 0 || tmp < 0 && n > 0 {
			cnt++
		}
	}

	fmt.Print(cnt)
}
