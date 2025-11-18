package main

import "fmt"

func main() {
	arrayWeek := [7]string{"Понедельник",
		"Вторник",
		"Среда",
		"Четверг",
		"Пятница",
		"Суббота",
		"Воскресенье",
	}

	ost := arrayWeek[2:]

	for i := 0; i < 7; i++ {
		ost = append(ost, "Пятница")
	}

	fmt.Println(ost)
}
