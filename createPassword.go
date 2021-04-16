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

	//your code here

	if len(name) < 5 {
		fmt.Println("Error ! Anda harus memasuskan minimal 5 karakter ")
	} else {
		var revers string
		var result string
		var first string = string(name[0])
		var last string = string(name[len(name)-1])
		for i := len(name) - 1; i >= 0; i-- {
			revers += string(name[i])
		}
		result = strings.ToUpper(first) + revers + strings.ToUpper(last) + strconv.Itoa(len(name))
		fmt.Println(result)
	}
}
