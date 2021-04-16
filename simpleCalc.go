package main

import (
	"fmt"
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
	var newResult int // result for modulus operation
	switch method {
	case "+":
		result = num1 + num2
		fmt.Printf("result: %.2f\n", result)
	case "-":
		result = num1 - num2
		fmt.Printf("result: %.2f\n", result)
	case "/":
		result = num1 / num2
		fmt.Printf("result: %.2f\n", result)
	case "*":
		result = num1 * num2
		fmt.Printf("result: %.2f\n", result)
	case "%":
		newNum1 := int(num1)
		newNum2 := int(num2)
		newResult = newNum1 % newNum2
		fmt.Printf("result: %d\n", newResult)
	default:
		fmt.Println("error operasi yang dimasukkan salah!")
	}
}
