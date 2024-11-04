// Данная задача ориентирована на реальную работу с данными в формате JSON.
// Для решения мы будем использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности),
// который был размещен на web-портале data.gov.ru.
//
// Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json,
// а сами данные, которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json. Файлы размещены в нашем репозитории на github.com.
//
// Для того, чтобы показать, что вы действительно смогли декодировать документ
// вам необходимо в качестве ответа записать сумму полей global_id всех элементов, закодированных в файле.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type MyStruct struct {
	GlobalID int `json:"global_id"`
}

func main() {
	jsonFile, err := os.Open("./data-20190514T0100.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer jsonFile.Close()

	var items []MyStruct

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&items)
	if err != nil {
		log.Fatal(err)
	}

	var total int

	for _, item := range items {
		total += item.GlobalID
	}

	fmt.Printf("Сумма global_id: %v\n", total)

}
