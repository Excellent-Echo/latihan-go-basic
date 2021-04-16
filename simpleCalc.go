package main

import (
	"fmt"
)

func main() {
	var (
		num1, num2, hasil int32
		method            string
	)
	fmt.Print("input calculator : ")
	fmt.Scanf("%d %v %d", &num1, &method, &num2)

	//your code here
	switch method {
	case "+":
		hasil = num1 + num2
		fmt.Printf("Hasil: %d", hasil)
	case "-":
		hasil = num1 - num2
		fmt.Printf("Hasil: %d", hasil)
	case "/":
		hasil = num1 / num2
		fmt.Printf("Hasil: %d", hasil)
	case "X":
		hasil = num1 * num2
		fmt.Printf("Hasil: %d", hasil)
	default:
		fmt.Println("input salah")
	}

}
