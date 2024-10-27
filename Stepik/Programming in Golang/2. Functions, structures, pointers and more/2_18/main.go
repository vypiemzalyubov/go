// Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
//
// Входные данные
// Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.
//
// Выходные данные
// Выведите максимальную цифру, которая встречается во введенной строке.

package main

import (
	"errors"
	"fmt"
)

func main() {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("wrong input type")
	}

	res, err := findMax(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func findMax(s string) (string, error) {
	res := '0'

	if len(s) < 0 {
		return "", errors.New("error")
	}

	for _, v := range s {
		if v > res {
			res = v
		}
	}

	return string(res), nil
}
