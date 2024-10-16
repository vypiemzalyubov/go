// Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр,
// на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.
// Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .
// По данному числу определите его цифровой корень.
//
// Входные данные
// Вводится одно натуральное число n, не превышающее 10**7.
//
// Выходные данные
// Вывести цифровой корень числа n.

package main

import "fmt"

func main() {
	var n, count, tmp int
	fmt.Scan(&n)

	for n > 0 {
		count += n % 10
		n /= 10
	}

	if count < 10 {
		fmt.Println(count)
	} else {
		for count > 0 {
			tmp += count % 10
			count /= 10
		}
		fmt.Println(tmp)
	}
}
