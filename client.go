package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Отправляем GET-запрос к нашему серверу
	resp, err := http.Get("http://localhost:8080/joke")
	if err != nil {
		fmt.Println("Ошибка при запросе:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Printf("Шутка от сервера: %s\n", string(body))
}
