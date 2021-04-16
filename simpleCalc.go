package main

import "fmt"

func main() {
	var (
		num1, num2 int
		method     string
	)
	fmt.Print("input calculator : ")
	fmt.Scanf("%d %s %d", &num1, &method, &num2)

	//your code here
	//state := true

	//for state == true {
	if method == "+" {
		fmt.Println(num1 + num2)
	} else if method == "-" {
		fmt.Println(num1 - num2)
	} else if method == "/" {
		fmt.Println(num1 / num2)
	} else if method == "%" {
		fmt.Println(num1 % num2)
	} //else {
	//	fmt.Println("error operasi yang dimasukan salah")
	//}
	//}
}
