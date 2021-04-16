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
		fmt.Println("error! name kurang dari 5")
	} else {
		var dibalik string
		var result string
		var awal string = string(name[0])
		var akhir string = string(name[len(name)-1])
		for i := len(name) - 1; i >= 0; i-- {
			dibalik += string(name[i])
		}
		result = strings.ToUpper(awal) + dibalik + strings.ToUpper(akhir) + strconv.Itoa(len(name))
		fmt.Println(result)
	}
}
