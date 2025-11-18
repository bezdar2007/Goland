package main

import "fmt"

func divmod(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {

	var a, b int
	fmt.Println("Введите два числа через пробел: ")
	fmt.Scan(&a)
	fmt.Scan(&b)

	if b != 0 {
		quotient, remainder := divmod(a, b)
		fmt.Printf("Частное равно: %d\n", quotient)
		fmt.Printf("Остаток равен: %d", remainder)
	} else {
		fmt.Println("division by zero")
	}
}
