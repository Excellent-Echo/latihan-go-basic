package main

import "fmt"

func main() {
	var (
		num1, num2 int
		method     string
	)
	fmt.Print("input calculator : ")
	fmt.Scanf("%d %v %d", &num1, &method, &num2)

	switch method {
	case "+":
		fmt.Println(num1, " + ", num2, " = ", num1+num2)
	case "-":
		fmt.Println(num1, " - ", num2, " = ", num1-num2)
	case "*":
		fmt.Println(num1, " * ", num2, " = ", num1*num2)
	case "/":
		fmt.Println(num1, " / ", num2, " = ", num1/num2)
	case "%":
		fmt.Println(num1, " % ", num2, " = ", (num1 % num2))
	case "default":
		fmt.Println("Inputan anda salah, silakan cek kembali!")
	}

}
