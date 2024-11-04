// На стандартный ввод подаются данные о студентах университетской группы в формате JSON:
//
// {
//     "ID":134,
//     "Number":"ИЛМ-1274",
//     "Year":2,
//     "Students":[
//         {
//             "LastName":"Вещий",
//             "FirstName":"Лифон",
//             "MiddleName":"Вениаминович",
//             "Birthday":"4апреля1970года",
//             "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//             "Phone":"+7(948)709-47-24",
//             "Rating":[1,2,3]
//         },
//         {
//             // ...
//         }
//     ]
// }
//
// В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
// Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы.
// Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:
// {
//     "Average": 14.1
// }
//
// Как вы понимаете, для декодирования используется функция Unmarshal, а для кодирования MarshalIndent (префикс - пустая строка, отступ - 4 пробела).

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	LastName   string `json:"LastName"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	Birthday   string `json:"Birthday"`
	Address    string `json:"Address"`
	Phone      string `json:"Phone"`
	Rating     []int  `json:"Rating"`
}

type Group struct {
	ID       int       `json:"ID"`
	Number   string    `json:"Number"`
	Year     int       `json:"Year"`
	Students []Student `json:"Students"`
}

type Result struct {
	Average float64 `json:"Average"`
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Ошибка чтения данных:", err)
		return
	}

	var group Group
	err = json.Unmarshal(data, &group)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return
	}

	totalRatings := 0
	for _, student := range group.Students {
		totalRatings += len(student.Rating)
	}

	average := 0.0
	if len(group.Students) > 0 {
		average = float64(totalRatings) / float64(len(group.Students))
	}

	result := Result{Average: average}

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
		return
	}

	fmt.Println(string(output))
}
