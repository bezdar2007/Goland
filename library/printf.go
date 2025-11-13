package main

import "fmt"

func main() {

	var a int = 100
	//Форматированный вывод (Форматированная строка, последующее значение подставляемое в формат (
	//%d - десятичных чисел)
	//%v - вещественных чисел
	//%s - строки
	//%g -
	//%T - герб - тип значения)
	fmt.Printf("a = %d, тип переменной: (type: %T)\n", a, a)

	b := 3.14
	fmt.Printf("b = %v, тип переменной: (type: %T)\n", b, b)

	var x string = "Текст"
	fmt.Printf("x = %s, тип переменной: (type: %T)\n", x, x)

	//Объявление нескольких переменных
	var num1, stroka = 1, "какой то текст"
	fmt.Printf("num1 = %d, stroka = %s", num1, stroka)

	const pi float64 = 3.1415
	fmt.Println("Число Pi равно", pi)
}
