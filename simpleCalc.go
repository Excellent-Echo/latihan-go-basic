package main

import "fmt"

func main() {
	var (
		num1, num2 int
		method     string
	)
	fmt.Print("input calculator : ")
	fmt.Scanf("%b %v %b", &num1, &method, &num2)

	if method == "+" {
		fmt.Println(num1 + num2)
	} else if method == "-" {
		fmt.Println(num1 - num2)
	} else if method == "/" {
		fmt.Println(num1 / num2)
	} else if method == "%" {
		fmt.Println(int(num1) % int(num2))
	} else {
		fmt.Println("error operasi yang dimasukkan salah!")
	}
}
