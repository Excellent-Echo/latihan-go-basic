package main

import "fmt"

func main() {
	var (
		num1, num2 float32
		method     string
	)
	fmt.Print("input calculator : ")
	fmt.Scanf("%b %v %b", &num1, &method, &num2)

	//your code here

	switch method {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	case "%":
		fmt.Println(int(num1) % int(num2))
	default:
		fmt.Println("error operator yang dimasukkan salah!")
	}
}
