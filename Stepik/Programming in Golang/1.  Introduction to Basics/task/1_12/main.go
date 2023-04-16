// 1.12 По данному трехзначному числу определите, все ли его цифры различны.
//      На вход подается одно натуральное трехзначное число.
//      Выведите "YES", если все цифры числа различны, в противном случае - "NO".

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	third_digit := n % 10
	second_digit := n / 10 % 10
	first_digit := n / 100

	if first_digit != second_digit && first_digit != third_digit && second_digit != third_digit {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
