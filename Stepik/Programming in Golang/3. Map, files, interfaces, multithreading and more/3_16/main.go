// На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
// Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	inputedTime, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	splitedTime := strings.Split(inputedTime, ",")
	splitedTime[1] = strings.TrimRight(splitedTime[1], "\n")

	layout := "02.01.2006 15:04:05"
	parseTime1, err := time.Parse(layout, splitedTime[0])
	if err != nil {
		panic(err)
	}
	parseTime2, err := time.Parse(layout, splitedTime[1])
	if err != nil {
		panic(err)
	}

	var result time.Duration

	if parseTime1.Before(parseTime2) {
		result = time.Since(parseTime1) - time.Since(parseTime2)
	} else {
		result = time.Since(parseTime2) - time.Since(parseTime1)
	}

	fmt.Println(result.Round(time.Second))
}
