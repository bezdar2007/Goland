package main

import "fmt"

func main() {

	var a [4]int
	a[0] = 1
	i := a[0]
	fmt.Println(i)

	b := [2]string{"Penn", "Teller"}
	fmt.Println(b)
	b1 := [...]string{"Penn", "Teller"}
	fmt.Println(b1)

}
