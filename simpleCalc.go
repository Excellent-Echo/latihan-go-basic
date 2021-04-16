//package main

import "fmt"

func main() {
	var (
		num1, num2 int
		method     string
	)
	fmt.Print("input calculator : ")
	fmt.Scanf("%d %v %d", &num1, &method, &num2)

	//your code here

	switch method {
	case "+":
		fmt.Println("hasil: ", num1+num2)
	case "-":
		fmt.Println("hasil: ", num1-num2)
	case "%":
		fmt.Println("hasil: ", num1%num2)
	case "*":
		fmt.Println("hasil: ", num1*num2)
	case "/":
		fmt.Println("hasil: ", num1/num2)
	default:
		fmt.Println("error operasi yang dimasukkan salah!!!!!!!!!!!")
	}

}
