// Напишите веб сервер, который по пути /get отдает текст "Hello, web!".
// Порт должен быть :8080.

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, web!"))
}

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
