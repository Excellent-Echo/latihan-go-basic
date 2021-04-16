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
	nameLength := len(name)
	password := ""

	if nameLength >= 5 {
		for i := nameLength - 1; i >= 0; i-- {
			password += string(name[i])
		}
		start := strings.ToUpper(string(name[0]))
		finish := strings.ToUpper(string(name[nameLength-1]))
		results := start + password + finish + strconv.Itoa(nameLength)
		fmt.Println("Password = ", results)
	} else {
		fmt.Println("Panjang Nama yang anda input terlalu pendek")
	}
}
