package main

import "fmt"

func main() {

	//типы данных
	var a int8 = -1
	var b uint8 = 2
	var c byte = 3
	var d rune = -7
	var f uint32 = 8
	var g uint64 = 10
	var h int = 192
	var j uint = 10
	var e float32 = 19
	var i float32 = 4.6
	var k float64 = 0.3232

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)
	fmt.Println("h: ", h)
	fmt.Println("j: ", j)
	fmt.Println("f: ", e)
	fmt.Println("g: ", i)
	fmt.Println("d: ", k)

	var compl complex64 = 1 + 2i
	var compl1 complex128 = 4 + 3i

	fmt.Print(compl, compl1)
}
