// Представьте что вы работаете в большой компании где используется модульная архитектура.
// Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные.
// Вы же пишете функцию которая считывает две переменных типа string ,  а возвращает число типа int64 которое нужно получить сложением этих строк.
//
// Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше.
// Поэтому он решил подшутить над вами и подсунул вам свинью. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.
//
// Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние символы и спец знаки.
// Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes. Они уже импортированы, вам ничего импортировать не нужно!
//
// Считывать и выводить ничего не нужно!
// Ваша функция должна называться adding()!
// Считайте что функция и пакет main уже объявлены!

package main

import "strconv"

func main() {
	adding("%^80", "hhhhh20&&&&nd")
}

func adding(a, b string) int64 {
	var d1, d2 int

	result1, err := strconv.Atoi(a)
	if err != nil {
		tmp1 := ""
		for _, v := range a {
			if v > 47 && v < 58 {
				tmp1 += string(v)
			}
		}
		r1, _ := strconv.Atoi(string(tmp1))
		d1 = r1
	} else {
		d1 = result1
	}

	result2, err := strconv.Atoi(b)
	if err != nil {
		tmp2 := ""
		for _, v := range b {
			if v > 47 && v < 58 {
				tmp2 += string(v)
			}
		}
		r2, _ := strconv.Atoi(string(tmp2))
		d2 = r2
	} else {
		d2 = result2
	}

	return int64(d1 + d2)
}