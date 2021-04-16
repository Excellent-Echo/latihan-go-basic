package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		num1, num2 float32
		method     string
	)
	fmt.Print("input calculator : ")
	fmt.Scanf("%b %v %b", &num1, &method, &num2)

	//your code here
	var result float32
	switch method {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "/":
		result = num1 / num2
	case "*":
		result = num1 * num2
	case "%":
		result = math.Mod(num1, num2)
	default:
		fmt.Println("error operasi yang dimasukkan salah!")
	}
	fmt.Printf("result: %f\n", result)
}
