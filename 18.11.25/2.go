package main

import (
	"fmt"
)

func main() {

	arrayWeek := [7]string{"Понедельник",
		"Вторник",
		"Среда",
		"Четверг",
		"Пятница",
		"Суббота",
		"Воскресенье",
	}
	today := arrayWeek[1]
	ost := arrayWeek[2:]

	fmt.Printf("Все дни недели: %v\n", arrayWeek)
	fmt.Printf("Сегодня: %v\n", today)
	fmt.Printf("Остались: %v\n", ost)

}
