// На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var a, b float64
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("wrong input type")
	}

	res, err := findG(a, b)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(res)
}

func findG(a, b float64) (float64, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("error")
	}
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2)), nil
}
