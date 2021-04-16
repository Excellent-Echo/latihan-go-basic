package main

import "fmt"

func main() {
	var (
		num1, num2 int32
		method     string
	)
	fmt.Print("input calculator : ")
	fmt.Scanf("%d %v %d", &num1, &method, &num2)

	//your code here
	switch method {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "/":
		fmt.Println(num1 / num2)
	case "*":
		fmt.Println(num1 * num2)
	case "%":
		fmt.Println(num1 % num2)
	default:
		fmt.Println("Error. Invalid input !!")
	}
}
