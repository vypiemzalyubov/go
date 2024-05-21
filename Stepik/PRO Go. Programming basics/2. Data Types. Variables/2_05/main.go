// Напишите программу, которая считывает три строки по очереди, а затем выводит их в той же последовательности, каждую на отдельной строчке.
//
// Формат входных данных
// На вход программе подаются три строки, каждая на отдельной строке.
//
// Формат выходных данных
// Программа должна вывести введенные строки в той же последовательности, каждую на отдельной строке.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	_ = scanner.Scan()
	s1 := scanner.Text()

	_ = scanner.Scan()
	s2 := scanner.Text()

	_ = scanner.Scan()
	s3 := scanner.Text()

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
