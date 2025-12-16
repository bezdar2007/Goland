package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем обработчик всех запросов к пути "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go!")
	})

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на порту 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
