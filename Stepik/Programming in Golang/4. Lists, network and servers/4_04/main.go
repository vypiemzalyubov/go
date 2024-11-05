// Сделайте HTTP запрос на сервер по пути http://127.0.0.1:5555/get и напечатайте ответ сервера (только тело).

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:5555/get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data)
}
