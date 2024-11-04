// Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап конвейера
// только если оно отличается от того, что пришло ранее.
//
// Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки,
// во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения, которые не повторяются подряд.
// Не забудьте закрыть канал ;)
//
// Функция должна называться removeDuplicates()

package main

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go removeDuplicates(c1, c2)
}

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream)

	var lastValue string

	for value := range inputStream {
		if value != lastValue {
			outputStream <- value
			lastValue = value
		}
	}
}
