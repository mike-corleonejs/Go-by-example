package main

import "fmt"

func main() {
	// string
	var str = "initial"
	fmt.Println(str)

	// declare multiple
	var number1, number2 = 1, 2
	fmt.Println(number1, number2)

	// boolean
	var isValid = true
	fmt.Println(isValid)

	// type of declare varibale, default varibale equal to 0
	var initInt int
	fmt.Println(initInt)

	// short declare, only use inside func
	Str := "Cú pháp khai báo nhanh"
	fmt.Println(Str)
}
