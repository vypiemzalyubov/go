// Используем анонимные функции на практике.
//
// Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint,
// которую внутри функции main в дальнейшем можно будет вызвать по имени fn.
// Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0.
// Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.
//
// Пакет main объявлять не нужно!
// Вводить и выводить что-либо не нужно!
// Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := uint(727178)

	fn := func(n uint) uint {
		str := ""
		strN := strconv.FormatUint(uint64(n), 10)

		for _, v := range strN {
			digit, _ := strconv.Atoi(string(v))
			if digit%2 == 0 && digit != 0 {
				str += string(v)
			}
		}

		num, _ := strconv.ParseUint(str, 10, 64)
		if num == 0 {
			return 100
		}
		return uint(num)

	}

	res := fn(n)
	fmt.Println(res)
}