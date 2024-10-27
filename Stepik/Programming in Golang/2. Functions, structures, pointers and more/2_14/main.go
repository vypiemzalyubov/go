// Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он может содержать только арабские цифры и буквы латинского алфавита.
// На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var s string
	fmt.Scan(&s)
	result := "Ok"

	if utf8.RuneCountInString(s) < 5 {
		result = "Wrong password"
	} else {
		for _, v := range s {
			if unicode.IsDigit(v) || unicode.Is(unicode.Latin, v) {
				continue
			} else {
				result = "Wrong password"
			}
		}
	}

	fmt.Println(result)
}
