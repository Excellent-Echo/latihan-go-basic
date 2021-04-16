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
	nameLen := len(name)
	pass := ""
	if nameLen <= 5 {
		fmt.Println("Error, nama minimal 5 karakter")
	} else {
		for i := nameLen - 1; i >= 0; i-- {
			pass += string(name[i])
		}
		awal := strings.ToUpper(string(name[0]))
		akhir := strings.ToUpper(string(name[nameLen-1]))
		hasil := awal + pass + akhir + strconv.Itoa(nameLen)
		fmt.Println(hasil)
	}
}
