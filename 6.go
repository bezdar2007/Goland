package main

import "fmt"

func main() {
	var num int = 0

	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num > 0 {
		if num%2 == 0 {
			fmt.Println("positive even")
		} else {
			fmt.Println("negative even")
		}
	} else if num < 0 {
		if num%2 != 0 {
			fmt.Println("positive odd")
		} else {
			fmt.Println("negative odd")
		}
	} else {
		fmt.Println("zero")
	}
}
