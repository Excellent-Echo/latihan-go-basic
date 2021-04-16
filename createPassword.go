package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		name, result string
	)

	fmt.Print("masukkan nama untuk diubah menjadi password : ")
	fmt.Scanf("%s", &name)

	//your code here
	if len(name) < 5 {
		fmt.Println("password harus lebih dari 5 karakter")
	} else {
		for i := len(name) - 1; i >= 0; i-- {
			if i == len(name)-1 {
				result += strings.ToUpper(string(name[0])) + strings.ToLower(string(name[i]))
			} else if i == 0 {
				result += strings.ToLower(string(name[i])) + strings.ToUpper(string(name[len(name)-1]))
			} else {
				result += strings.ToLower(string(name[i]))
			}
		}
	}
	fmt.Printf("Result: %s", result+strconv.Itoa(len(name)))
}

// masukkan nama untuk diubah menjadi password : tarjo
// Result: TojratO5
