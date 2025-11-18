package main

import "fmt"

func main() {
	a := 10
	b := 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	if b != 0 {
		fmt.Printf("%.4f\n", float64(a)/float64(b))
	}
}
