package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		name, tmp string
	)

	fmt.Print("masukkan nama untuk diubah menjadi password : ")
	fmt.Scanf("%s", &name)

	//your code here
	length := len(name)
	if length < 5 {
		fmt.Println("Password must be 5 or more")
	} else {
		//Pembalik kata
		for i := length - 1; i >= 0; i-- {
			//fmt.Println(string(name[i]))
			tmp += string(name[i])
		}
		depan := strings.ToUpper(string(tmp[0]))
		belakang := strings.ToUpper(string(tmp[length-1]))
		res := depan + tmp + belakang + strconv.Itoa(length)
		fmt.Println(res)

	}
}
