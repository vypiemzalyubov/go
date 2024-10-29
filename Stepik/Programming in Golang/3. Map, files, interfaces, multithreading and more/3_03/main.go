// Напишите функцию с именем convert, которая конвертирует входное число типа int64 в uint16.
//
// Считывать и выводить ничего не нужно!
// Считайте что функция main и пакет main уже объявлены!

package main

func main() {
	var n int64
	convert(n)
}

func convert(n int64) uint16 {
	return uint16(n)
}
