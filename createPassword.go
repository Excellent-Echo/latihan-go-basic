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
		fmt.Println("nama minimal 5 karakter")
		return
	}

	firstLetter := strings.ToUpper(string(name[0]))
	lastLetter := strings.ToUpper(string(name[len(name)-1]))
	lastPswChar := strconv.Itoa(len(name))

	var reverseName string
	for i := len(name) - 1; i >= 0; i-- {
		reverseName += string(name[i])
	}

	fmt.Println("Your password:", firstLetter+reverseName+lastLetter+lastPswChar)
}
