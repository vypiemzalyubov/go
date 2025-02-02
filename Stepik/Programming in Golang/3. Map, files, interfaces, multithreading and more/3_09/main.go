// Ранее в рамках этого курса при решении задач требовалось прочитать что-то со стандартного ввода и вывести результат соответственно в стандартный вывод.
// При этом кто-то использовал пакет fmt, а кто-то - bufio + os. Все эти пакеты имеют свои особенности,
// поэтому в этой задаче мы попробуем решить знакомую нам проблему с помощью пакетов, которые кто-то мог до этого момента и не применять: bufio + os + strconv.
//
// Задача состоит в следующем: на стандартный ввод подаются целые числа в диапазоне 0-100, каждое число подается на стандартный ввод с новой строки
// (количество чисел не известно). Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.
//
// Несколько подсказок: для чтения вы можете использовать типы bufio.Reader и bufio.Scanner, а для записи - bufio.Writer.
// При чтении помните об ошибке io.EOF. Конвертирование данных из строки в целое число и обратно осуществляется функциями Atoi и Itoa из пакета strconv
// соответственно. Чтение производится из стандартного ввода (os.Stdin), а запись - в стандартный вывод (os.Stdout).
//
// Все указанные в тексте задачи пакеты (strconv, bufio, os, io), уже импортированы (другие использовать нельзя), package main объявлен.

package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	result := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		result += num
	}

	str := strconv.Itoa(result)
	io.WriteString(os.Stdout, str)
}
