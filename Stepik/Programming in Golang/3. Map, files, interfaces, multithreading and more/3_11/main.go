// Данная задача в основном ориентирована на изучение типа bufio.Reader, поскольку этот тип позволяет считывать данные постепенно.
//
// В тестовом файле, который вы можете скачать из нашего репозитория на github.com, содержится длинный ряд чисел, разделенных символом ";".
// Требуется найти, на какой позиции находится число 0 и указать её в качестве ответа.
// Требуется вывести именно позицию числа, а не индекс (то-есть порядковый номер, нумерация с 1).
//
// Например:  12;234;6;0;78 :
// Правильный ответ будет порядковый номер числа: 4

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	csvFile, err := os.Open("./task.data")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer csvFile.Close()

	rd := bufio.NewReader(csvFile)
	reader := csv.NewReader(rd)
	reader.Comma = ';'

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Ошибка при чтении файла:", err)
			return
		}

		for num, item := range record {
			if item == "0" {
				fmt.Println(num + 1)
				break
			}
		}
	}
}
