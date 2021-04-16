package main

import "fmt"

func main() {
	var (
		num1, num2 float32
		method     string
	)
	fmt.Print("input calculator : ")
	fmt.Scanf("%d %v %d", &num1, &method, &num2)

	//your code here

	switch method{
	case 1:
		fmt.Println("hasil: ", num1+num2)

}
