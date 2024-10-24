// На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	r := []rune(text)

	if unicode.IsUpper(r[0]) && r[len(r)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
