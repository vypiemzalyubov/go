// Напишите программу для пересчёта величины временного интервала, заданного в минутах, в величину, выраженную в часах и минутах.
//
// Формат входных данных
// На вход программе подаётся целое число – количество минут.
//
// Формат выходных данных
// Программа должна вывести текст в соответствии с условием задачи.

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%d мин - это %d час %d минут.", n, n/60, n%60)
}