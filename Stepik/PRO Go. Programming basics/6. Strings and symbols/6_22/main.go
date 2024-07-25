// Сейчас в городе Арторзе популярна троичная система счисления. Для телеграфной передачи чисел, записанных в троичной системе счисления, используется азбука Арторзе.
// Цифра 0 передается как ., 1 как -., 2 как --.
// Расшифровка кода Арторзе чисел — очень важная и ответственная работа. Ваша задача — расшифровать заданное в коде Арторзе троичное число.
//
// Входные данные
// В первой строке записано число в коде Арторзе. Длина кода не меньше 1 и не больше 200 символов.
// Гарантируется, что заданная строка — корректный код Арторзе некоторого числа в троичной системе счисления (число могло содержать лидирующие нули).
//
// Выходные данные
// Выведите расшифровку заданного кода Арторзе. Расшифрованное число может содержать лидирующие нули.

package main

import (
	"fmt"
)

func main() {
	var code string
	_, _ = fmt.Scan(&code)

	for i := 0; i < len(code); {
		switch {
		case i < len(code)-1 && code[i] == '-' && code[i+1] == '.':
			fmt.Print(1)
			i += 2
		case i < len(code)-1 && code[i] == '-' && code[i+1] == '-':
			fmt.Print(2)
			i += 2
		case code[i] == '.':
			fmt.Print(0)
			i++
		default:
			i++
		}
	}
}
