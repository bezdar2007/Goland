package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 5; j++ {
		println(j)
	}

	for k := range 6 {
		if k == 1 {
			continue
		} else {
			fmt.Println(k)
		}
	}

}
