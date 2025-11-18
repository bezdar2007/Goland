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
	lenArrayWeek := len(arrayWeek)
	capArrayWeek := cap(arrayWeek)

	slice := arrayWeek[2:]

	lenSlice := len(slice)
	capSlice := cap(slice)

	fmt.Printf("Длина массива: %v\n", lenArrayWeek)
	fmt.Printf("Ёмкость массива: %v\n", capArrayWeek)

	fmt.Printf("Длина среза: %v\n", lenSlice)
	fmt.Printf("Ёмкость среза: %v\n", capSlice)

}
