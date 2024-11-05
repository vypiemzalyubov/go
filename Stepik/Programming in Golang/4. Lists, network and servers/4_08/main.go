// Напиши веб сервер (порт :3333) - счетчик который будет обрабатывать GET (/count) и POST (/count) запросы:
// - GET:  возвращает счетчик
// - POST: увеличивает ваш счетчик на значение  (с ключом "count") которое вы получаете из формы,
//         но если пришло НЕ число то нужно ответить клиенту: "это не число" со статусом http.StatusBadRequest (400).

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(counter)))
	case http.MethodPost:
		err := r.ParseForm()
		if err == nil {
			countStr := r.FormValue("count")
			if countStr == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("это не число"))
				return
			}

			count, err := strconv.Atoi(countStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("это не число"))
				return
			}

			counter += count
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Метод не поддерживается"))
	}
}

func main() {
	http.HandleFunc("/count", handler)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
