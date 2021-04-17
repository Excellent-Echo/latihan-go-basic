package main

import "fmt"

func main() {
	var (
		num1, num2 int
		method string
	)
	fmt.Println("Kalkulator : ")
	fmt.Scanf("%d %v %d", &num1, &method, &num2)

	//your code here
	switch method {
	case "+":
		fmt.Printf("Hasil : %d", num1 + num2)
	case "-":
		fmt.Printf("Hasil : %d", num1 - num2)
	case "*":
		fmt.Printf("Hasil : %d", num1 * num2)
	case "/":
		fmt.Printf("Hasil : %d", num1 / num2)
	case "%":
		fmt.Printf("Hasil : %d", num1 % num2)
	default:
		fmt.Println("Masukan salah!")
	}
}
