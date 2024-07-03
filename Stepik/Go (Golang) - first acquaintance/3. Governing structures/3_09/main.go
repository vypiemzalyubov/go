// На вход подается целое число, сумма денег, которые у вас есть. Ваша задача - вывести марку телефона, которую вы можете себе позволить купить.
// - Если сумма больше 1000 - вывести Apple
// - Если сумма от 500 до 1000 (включительно) - вывести Samsung
// - Если сумма меньше 500 - вывести Nokia с фонариком

package main

import "fmt"

func main() {
	var sum int
	fmt.Scan(&sum)
	if sum > 1000 {
		fmt.Println("Apple")
	} else if sum >= 500 && sum <= 1000 {
		fmt.Println("Samsung")
	} else {
		fmt.Println("Nokia с фонариком")
	}
}
