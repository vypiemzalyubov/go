// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, result string
	fmt.Scan(&s)

	for _, v := range s {
		if strings.Count(s, string(v)) < 2 {
			result += string(v)
		}
	}

	fmt.Println(result)
}
