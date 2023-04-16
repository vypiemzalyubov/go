// 1.14 Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
//      На вход подается номер билета - одно шестизначное  число.
//      Выведите "YES", если билет счастливый, в противном случае - "NO".

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	first_d := n / 100000
	second_d := n / 10000 % 10
	third_d := n / 1000 % 10
	fourth_d := n / 100 % 10
	fifth_d := n / 10 % 10
	sixth_d := n % 10

	if first_d+second_d+third_d == fourth_d+fifth_d+sixth_d {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
