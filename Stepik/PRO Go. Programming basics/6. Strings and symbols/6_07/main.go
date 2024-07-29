// Определить, является ли введенное слово идентификатором, т.е. начинается ли оно с английской буквы в любом регистре или со знака подчеркивания,
// так же она НЕ должна содержать других символов, КРОМЕ букв английского алфавита (в любом регистре), цифр и знака подчеркивания.
//
// Входные данные
// На вход программе подается одна строка.
//
// Выходные данные
// Если строка является идентификатором, необходимо вывести "YES", иначе - "NO".

package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text string
	_, _ = fmt.Scan(&text)
	flag := "YES"
	re := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

	if !re.MatchString(text) {
		flag = "NO"
	}
	fmt.Print(flag)
}