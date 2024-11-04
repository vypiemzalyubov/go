// На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере).
// Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).
//
// Требуется считать данные о продолжительности периода, преобразовать их в тип Duration,
// а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.
//
// Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

package main

import (
	"fmt"
	"time"
)

func main() {
	const now = 1589570165

	var min, sec int
	fmt.Scanf("%d мин. %d", &min, &sec)

	duration := time.Duration(min)*time.Minute + time.Duration(sec)*time.Second

	baseTime := time.Unix(now, 0).UTC()
	newTime := baseTime.Add(duration)
	fmt.Println(newTime.Format(time.UnixDate))
}
