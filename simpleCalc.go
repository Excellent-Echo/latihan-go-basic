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
			result := num1 + num2
			fmt.Println("Hasil penjumlahan :", result)
		case "-":
			result := num1 - num2
			fmt.Println("Hasil pengurangan :", result)
		case "/":
			result := num1 / num2
			fmt.Println("Hasil pembagian :", result)
		case "*": 
			result := num1 * num2
			fmt.Println("Hasil perkalian :", result)
		case "%": 
			result := num1 % num2
			fmt.Println("Hasil modulus :", result)
		default:
			fmt.Println("error operasi yang dimasukkan salah!")
		}
}
