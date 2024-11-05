// Считайте с консоли две переменные, сначала name, затем age.
// Сделайте HTTP запрос на http://127.0.0.1:8080/hello передав name и age в query параметрах, затем напечатайте ответ сервера (только тело).

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	var name, age string
	fmt.Scan(&name, &age)

	baseURL := "http://127.0.0.1:8080/hello"
	params := url.Values{}
	params.Add("name", name)
	params.Add("age", age)

	fullURL := baseURL + "?" + params.Encode()

	resp, err := http.Get(fullURL)
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
