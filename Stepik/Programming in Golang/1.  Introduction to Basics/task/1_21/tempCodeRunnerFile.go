package main

import "fmt"

func main() {
	var sequence, number int
	fmt.Scan(&sequence)
	for i := 1; i <= sequence; i++ {
		fmt.Scan(&number)
		if number <= 10 {
			continue
		} else if number >= 100 {
			break
		} else {
			fmt.Println(number)
		}
	}
}