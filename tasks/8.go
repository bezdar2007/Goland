package main

import "fmt"

func main() {

	var num int
	fmt.Println("Введите число: ")
	fmt.Scan(&num)

	summ := 0
	for i := 0; i <= num; i++ {
		summ += i
	}

	fmt.Printf("Сумма чисел равна %d", summ)
}
