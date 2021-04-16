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
	str := []int32(name)
	var result []int32

	if len(str) < 5 {
		"password minimal 5 karekter"
	}
	fmt.Printf(len(str))

	// var result string
	// for i := length; i < 1; i-- {
	// 	result = name[length]
	// }

	// fmt.Println(result)
}
