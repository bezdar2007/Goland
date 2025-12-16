package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Статические файлы
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Обработчики
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/todo-add", addFormHandler)

	// API эндпоинты
	http.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodosHandler(w, r)
		case http.MethodPost:
			addTodoHandler(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/todos/delete", deleteTodoHandler)

	fmt.Println("Сервер TODO запущен на http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
