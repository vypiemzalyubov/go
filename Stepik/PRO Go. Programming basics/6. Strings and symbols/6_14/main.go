// Даны две строчные буквы латинского алфавита. Выведите все строчные буквы латинского алфавита в алфавитном порядке, которые находятся между данными символами, а также их самих.
//
// Входные данные
// Вводятся две строчные буквы латинского алфавита.
//
// Выходные данные
// Выведите на одной строке через пробел все строчные буквы латинского алфавита в алфавитном порядке, которые находятся между данными символами, а также их самих.
//
// Примечание: первая буква не обязательно меньше, чем вторая буква.

package main

import (
	"fmt"
)

func main() {
	var c1, c2 rune
	_, _ = fmt.Scanf("%c\n%c", &c1, &c2)

	if c1 < c2 {
		for i := c1; i <= c2; i++ {
			fmt.Print(string(i), " ")
		}
	} else {
		for i := c2; i <= c1; i++ {
			fmt.Print(string(i), " ")
		}
	}
}