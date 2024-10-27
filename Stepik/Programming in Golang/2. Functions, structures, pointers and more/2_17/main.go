// Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
//
// Входные данные
// Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.
//
// Выходные данные
// Вывести строку, которая получится после добавления символов '*'.

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

	res, err := replace(s)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(res)
}

func replace(s string) (string, error) {
	var res string

	if len(s) < 0 {
		return "", errors.New("error")
	}

	for i, v := range s {
		if i != 0 {
			res += fmt.Sprintf("*%v", string(v))
		} else {
			res += string(v)
		}
	}

	return res, nil
}
