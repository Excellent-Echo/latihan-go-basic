package main

import (
	"fmt"

)

func main() {
	var (
		name string
	)

	fmt.Print("masukkan nama untuk diubah menjadi password : ")
	fmt.Scanf("%s", &name)

	//your code here

	if len(name) < 5 {

		fmt.Println("error, password minimal 5 karakter!")
	}else{
		for i := len(name) - 1 ; i > 0 ; i--{
			fmt.Println(i)
		}
	}

}
