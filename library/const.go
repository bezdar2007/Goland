package main

import "fmt"

func main() {

	const (
		kb = 1 << (10 * iota)
		mb
		gb
	)
	fmt.Println(kb, mb, gb)
}
