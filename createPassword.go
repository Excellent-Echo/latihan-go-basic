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

	fmt.Print("Masukkan nama untuk diubah menjadi password : ")
	fmt.Scanf("%s", &name)
	
	//your code here

	if len(name) < 5 {
		fmt.Printf("Inputan string harus minimal 5")
	} else {
		var reverse string
		var result string
		firstString := string(name[0])
		lastString := string(name[len(name)-1])
		for willReverse := len(name) - 1; willReverse >= 0; willReverse-- {
			reverse += string(name[willReverse])
		}
		result = strings.ToUpper(firstString) + reverse + strings.ToUpper(lastString) + strconv.Itoa(len(name))
		fmt.Printf(result)
	}
}