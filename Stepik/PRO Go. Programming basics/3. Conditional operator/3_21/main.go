// Даны действительные числа a, b, c. Найдите все решения квадратного уравнения ax**2+bx+c=0.
//
// Входные данные
// Вводятся три действительных числа a,b,c, при этом a≠0.
//
// Выходные данные
// Выведите два действительных числа, каждое на отдельной строке, если уравнение имеет два корня (сначала меньший, далее больший),
// одно действительное число – при наличии одного корня. При отсутствии действительных корней ничего выводить не нужно.

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
	} else if discriminant == 0 {
		fmt.Print((float64(-b) + math.Sqrt(float64(discriminant))) / float64(2*a))
	} else {
		root1 := (float64(-b) - math.Sqrt(float64(discriminant))) / float64(2*a)
		root2 := (float64(-b) + math.Sqrt(float64(discriminant))) / float64(2*a)

		if root1 < root2 {
			fmt.Println(root1)
			fmt.Println(root2)
		} else {
			fmt.Println(root2)
			fmt.Println(root1)
		}
	}
}
