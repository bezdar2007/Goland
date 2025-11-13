package main

import "fmt"

func AddNumbers(n1 int, n2 int) (int, int) {
	var summ int = n1 + n2
	var razn int = n1 - n2
	return summ, razn
}

func main() {

	a, b := AddNumbers(1, 2)

	fmt.Print(a, b)

}
