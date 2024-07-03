// Напишите функцию isEven(), которая принимает в качестве аргумента одно целое число и возвращает true если оно четное и false в ином случае.

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}