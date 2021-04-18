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
	if len(name) > 5 {
		fmt.Println("Password kurang dari 5 karakter")
	} else {
		var (
			berubah    string
			indexsatu  string = string(name[0])
			indexakhir string = string(name[len(name)-1]))
			hasilakhir string 
		)
		for i := len(name) - 1; i >= 0; i-- {
			berubah += string(name[i])
		}
		hasilakhir = strings.ToUpper(indexsatu) + berubah + strings.ToLower(indexakhir) + strconv.Itoa(len(name))
		fmt.Printf(hasilakhir)

	}
}
