package main

import (
	"fmt"
)

func main() {
	var (
		num1, num2, results int
		method              string
	)
	fmt.Printf(method)
	fmt.Print("input calculator : ")
	fmt.Scanf("%d %v %d", &num1, &method, &num2)
	//your code here

	if method == "+" {
		results = num1 + num2
		fmt.Println("Hasil = ", results)
	} else if method == "-" {
		fmt.Println("Hasil = ", num1-num2)
	} else if method == "*" {
		fmt.Println("Hasil = ", num1*num2)
	} else if method == "/" {
		fmt.Println("Hasil = ", num1/num2)
	} else if method == "%" {
		fmt.Println("Hasil = ", num1%num2)
	} else {
		fmt.Println("Format yang anda masukkan salah")
	}

}
