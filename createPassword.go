package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var (
		name, newName string
	)

	fmt.Print("masukkan nama untuk diubah menjadi password : ")
	fmt.Scanf("%s", &name)

	//your code here
	length := len(name)
	if length < 5 {
		fmt.Println("nama minimal 5 karakter")
	} else {
		// fmt.Println("sudah 5 karakter")
		for i := length - 1; i >= 0; i-- {
			// fmt.Println(string(name[i]))
			newName += string(name[i])
			// fmt.Println(newName)
		}
		firstIndex := strings.ToUpper(string(newName[0]))
		lastIndex := strings.ToUpper(string(newName[length-1]))
		// fmt.Println("ini index pertama", firstIndex)
		// fmt.Println("ini index terakhir", lastIndex)
		password := firstIndex + newName + lastIndex + strconv.Itoa(length)
		fmt.Println(password)
	}

}
