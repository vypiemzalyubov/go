// Недавно Артур выписал на листике k2 цифр 2, k3 цифр 3, k5 цифр 5 и k6 цифр 6.
// Любимые числа Артура — 32 и 256. Поэтому он, конечно же, решил составить из имеющихся цифр свои любимые числа.
// При этом он хочет, чтобы сумма составленных чисел была как можно больше. Помогите Артуру найти эту сумму!
// Каждую цифру можно использовать не более одного раза, то есть в составленных Артуром числах должно быть не больше k2 цифр 2, k3 цифр 3 и так далее.
// Неиспользованные цифры в сумме не учитываются.
//
// Входные данные
// Даны четыре целых числа k2, k3, k5, k6 — количество цифр 2, 3, 5 и 6 соответственно (0 ≤ k2, k3, k5, k6 ≤ 5⋅10**6).
//
// Выходные данные
// Выведите максимальную сумму любимых чисел Артура, которые можно составить имеющихся цифр.

package main

import (
	"fmt"
)

func main() {
	var k2, k3, k5, k6 int
	fmt.Scan(&k2, &k3, &k5, &k6)

	minDigit256 := min(k2, k5, k6)
	digit256 := minDigit256 * 256

	minDigit32 := min(k2-minDigit256, k3)
	digit32 := minDigit32 * 32

	fmt.Print(digit256 + digit32)
}
