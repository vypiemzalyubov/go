// На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

package main

import (
	"fmt"
)

func main() {
	var s, result string
	fmt.Scan(&s)

	for i, v := range s {
		if i%2 != 0 {
			result += string(v)
		}
	}

	fmt.Println(result)
}
