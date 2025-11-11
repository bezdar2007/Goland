package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число: ")
	fmt.Scan(&num)

	switch {
	case 90 <= num && num <= 100:
		fmt.Println("A")
	case 80 <= num && num <= 89:
		fmt.Println("B")
	case 70 <= num && num <= 79:
		fmt.Println("C")
	case 60 <= num && num <= 69:
		fmt.Println("D")
	case 0 <= num && num <= 59:
		fmt.Println("F")
	default:
		fmt.Println("invalid")
	}
}
