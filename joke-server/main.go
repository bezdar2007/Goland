package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var jokes = []string{
	"Почему программисты путают Хэллоуин и Рождество? Потому что Oct 31 == Dec 25.",
	"Алгоритм шуток: сначала сделай что-то смешное, потом оптимизируй.",
	"Программист не боится темноты — он боится пустого Stack Overflow.",
	"— Ты где? — Я в Go. — В каком? — В конкурентном.",
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	joke := jokes[rand.Intn(len(jokes))]

	tmpl, err := template.ParseFiles("templates/joke.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, joke)
	if err != nil {
		http.Error(w, "Ошибка рендера", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/joke", jokeHandler)

	http.ListenAndServe(":8080", nil)
}
