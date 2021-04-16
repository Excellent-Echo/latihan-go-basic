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

		fmt.Println("error, password minimal terdiri dari 5 karakter!")
	}else{
		var reserve string
		for i := len(name) - 1 ; i > 0 ; i--{
			reserve += string(name[i])
		
		}

		var result string
		var awal string= string(name[0])
		var akhir string=string(name[len(name)-1])

	
		result = strings.ToUpper(awal) + reserve + strings.ToUpper(akhir) + strconv.Itoa(len(name))
		fmt.Println(result)

	}

}
