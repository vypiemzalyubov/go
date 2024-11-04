// На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:
// 2020-05-15 08:00:00
//
// Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же формате.
// Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день, а затем вывести на стандартный вывод в том же формате.

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	parseTime, err := time.Parse("2006-01-02 15:04:05", scanner.Text())
	if err != nil {
		panic(err)
	}

	if parseTime.Hour() >= 13 {
		parseTime = parseTime.Add(24 * time.Hour)
	}

	fmt.Println(parseTime.Format("2006-01-02 15:04:05"))
}
