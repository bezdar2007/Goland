package main

import "fmt"

func main() {
	x := 5

	x += 3
	x += 2
	x *= 4
	x -= 10
	fmt.Println(x)

	var num int
	fmt.Print("Введите число: ")
	fmt.Scanln(&num)

	num += 3
	fmt.Println(num)

	num += 2
	fmt.Println(num)

	num *= 4
	fmt.Println(num)

	num -= 10
	fmt.Println(num)

}
