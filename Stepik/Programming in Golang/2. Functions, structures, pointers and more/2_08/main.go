// В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной жизни.
// Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.
//
// Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
// У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
//
// Если значение On == false, то оба метода вернут false.
//
// Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
//
// Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры с именем testStruct в функции main, в дальнейшем программа проверит результат.
//
// Закрывающая фигурная скобка в конце main() вам не видна, но она есть.
//
// Пакет main объявлять не нужно!
//
// Удачи!

package main

import "fmt"

type Task struct {
	On          bool
	Ammo, Power int
}

func (t *Task) Shoot() bool {
	if t.On == false {
		return false
	}
	if t.Ammo > 0 {
		t.Ammo--
		return true
	}
	return false
}

func (t *Task) RideBike() bool {
	if t.On == false {
		return false
	}
	if t.Power > 0 {
		t.Power--
		return true
	}
	return false
}

func main() {
	testStruct1 := Task{On: true, Ammo: 3, Power: 2}

	fmt.Println(testStruct1.Ammo)
	fmt.Println(testStruct1.Shoot())
	fmt.Println(testStruct1.Ammo)

	fmt.Println(testStruct1.Power)
	fmt.Println(testStruct1.RideBike())
	fmt.Println(testStruct1.Power)

	testStruct2 := Task{On: false, Ammo: 5, Power: 6}
	fmt.Println(testStruct2.Ammo)
	fmt.Println(testStruct2.Shoot())
	fmt.Println(testStruct2.Ammo)

	fmt.Println(testStruct2.Power)
	fmt.Println(testStruct2.RideBike())
	fmt.Println(testStruct2.Power)

}
