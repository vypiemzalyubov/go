// Определите являются ли билеты счастливыми. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
//
// Формат входных данных
// На вход подаются номера билетов - два шестизначных числа.
//
// Формат выходных данных
// Выведите 1, если они оба счастливые, в противном случае -1.
//
// Примечание: предполагается задачу решить с помощью функции.

package main

import (
	"fmt"
)

func main() {
	var ticket1, ticket2 string
	_, _ = fmt.Scan(&ticket1, &ticket2)

	if luckyTicket(ticket1)+luckyTicket(ticket2) == 2 {
		fmt.Print(1)
	} else {
		fmt.Print(-1)
	}
}

func luckyTicket(ticket string) int {
	if int(ticket[0]+ticket[1]+ticket[2]) == int(ticket[3]+ticket[4]+ticket[5]) {
		return 1
	} else {
		return 0
	}
}
