// На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
// Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var d string
	_, err := fmt.Scan(&d)
	if err != nil {
		fmt.Println("wrong input type")
	}

	res, err := findPow(d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func findPow(d string) (string, error) {
	res := ""

	if len(d) < 0 {
		return "0", errors.New("error")
	}

	for _, v := range d {
		f, _ := strconv.Atoi(string(v))
		res += strconv.Itoa(f * f)
	}

	return res, nil
}
