package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	jokes := []string{
		"Почему программисты путают Хэллоуин и Рождество? Потому что Oct 31 == Dec 25.",
		"Алгоритм шуток: сначала сделай что-то смешное, потом напиши код.",
		"— Программисту на свидании: Ты околок в тесте видела? — Нет... — Вот и правильно.",
	}
	// Выбираем случайную шутку из слайса
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(jokes))

	// Отправляем шутку в ответ (устанавливаем Content-Type text/plain)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, jokes[i])
}

func main() {
	http.HandleFunc("/joke", jokeHandler)
	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}
