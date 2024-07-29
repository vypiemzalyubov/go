// Вам необходимо написать функцию myPrint, которая может принимать на вход произвольное количество аргументов. Аргументы представляют собой целые числа.
// Далее все эти числа должны выводиться, каждое с новой строки.
//
// Написать нужно только функцию, которая должна выполнять то, что указано в задании.

func myPrint(x ...int) {
	for _, v := range x {
		fmt.Println(v)
	}
}