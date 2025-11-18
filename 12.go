package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c*9/5 + 32)
}

func main() {

	var value float64
	fmt.Print("Введите градусы: ")
	fmt.Scan(&value)

	C := Celsius(value)
	F := C.ToFahrenheit()
	res := fmt.Sprintf("%.2f", F)
	fmt.Printf("Значение в Фаренгейтах: %v", res)

}
