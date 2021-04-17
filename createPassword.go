package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		name string
		pw   string
	)

	fmt.Print("masukkan nama untuk diubah menjadi password : ")
	fmt.Scanf("%s", &name)

	if len(name) <= 5 {
		fmt.Println("Nama harus lebih dari 5 karakter!")
	} else {

		for i := len(name) - 1; i >= 0; i-- {
			pw += string(name[i])
		}

		st := strings.ToUpper(string(name[0]))
		en := strings.ToUpper(string(name[len(name)-1]))

		hasil := st + pw + en + strconv.Itoa(len(name))

		fmt.Println(hasil)
	}

}
