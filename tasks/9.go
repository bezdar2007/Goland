package main

import "fmt"

func summ(a int, b int) int {
	return a + b
}
func main() {
	var a, b int

	fmt.Println("Введите числа через пробел: ")
	fmt.Scan(&a)
	fmt.Scan(&b)

	res := summ(a, b)

	fmt.Println(res)
}
