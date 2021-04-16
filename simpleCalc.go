package main

import "fmt"

func main() {
	var (
		num1, num2, result int32
		method             string
	)
	fmt.Print("input calculator : ")
	fmt.Scanf("%d %v %d", &num1, &method, &num2)

	// your code here
	switch method {
	case "+":
		result = num1 + num2
		fmt.Printf("Hasil penjumlahan: %d", result)
	case "-":
		result = num1 - num2
		fmt.Printf("Hasil pengurangan: %d", result)
	case "x":
		result = num1 * num2
		fmt.Printf("Hasil perkalian: %d", result)
	case "/":
		result = num1 / num2
		fmt.Printf("Hasil Pembagian: %d", result)
	default:
		fmt.Println("Inputan salah")
	}
}
