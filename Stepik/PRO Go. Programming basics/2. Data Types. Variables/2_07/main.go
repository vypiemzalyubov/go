// Напишите программу, которая считывает строку-разделитель и три строки, а затем выводит указанные строки через разделитель.
//
// Формат входных данных
// На вход программе подаётся строка-разделитель и три строки, каждая на отдельной строке.
//
// Формат выходных данных
// Программа должна вывести введённые три строки через разделитель.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	_ = scanner.Scan()
	sep := scanner.Text()

	_ = scanner.Scan()
	s1 := scanner.Text()

	_ = scanner.Scan()
	s2 := scanner.Text()

	_ = scanner.Scan()
	s3 := scanner.Text()

	fmt.Println(s1 + sep + s2 + sep + s3)
}
