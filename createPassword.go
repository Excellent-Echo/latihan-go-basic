package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		name string
	)

	fmt.Print("masukkan nama untuk diubah menjadi password : ")
	fmt.Scanf("%s", &name)

	if len(name) >= 5 {
		var reverse string
		var result string
		var depan string = string(name[0])
		var belakang string = string(name[len(name)-1])

		for i := len(name) - 1; i >= 0; i-- {
			reverse += string(name[i])
		}
		result = strings.ToUpper(depan) + reverse + strings.ToUpper(belakang) + strconv.Itoa(len(name))
		fmt.Println(result)
	} else {
		fmt.Println("Error")
	}
}
