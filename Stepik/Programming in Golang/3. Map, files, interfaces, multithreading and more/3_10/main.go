// Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath,
// хотя для решения может быть использован также пакет archive/zip (поскольку файл с заданием предоставляется именно в этом формате).
//
// В тестовом архиве, который вы можете скачать из нашего репозитория на github.com, содержится набор папок и файлов.
// Один из этих файлов является файлом с данными в формате CSV, прочие же файлы структурированных данных не содержат.
//
// Требуется найти и прочитать этот единственный файл со структурированными данными (это таблица 10х10, разделителем является запятая),
// а в качестве ответа необходимо указать число, находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).

package main

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func walkFunc(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}

	file, _ := os.Open(path)
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if len(records) > 1 {
		fmt.Println(records[4][2])
		fmt.Println(info.Name())
		fmt.Println(path)
	}
	return nil
}

func main() {

	filepath.Walk("./task", walkFunc)

}
