// Напишите программу, которая считывает три строки по очереди, а затем выводит их в обратной последовательности, каждую на отдельной строчке.
//
// Формат входных данных
// На вход программе подается три строки, каждая на отдельной строке.
//
// Формат выходных данных
// Программа должна вывести введенные строки в обратной последовательности, каждую на отдельной строке.
//
// Примечания:
// Используйте 3 переменные для сохранения введённых строк текста.

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

	fmt.Println(s3 + "\n" + s2 + "\n" + s1)
}
