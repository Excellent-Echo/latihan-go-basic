package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		name, reverse string
	)

	fmt.Printf("Enter a name to be changed to a password (min 5 chars): ")
	fmt.Scanf("%s", &name)

	//your code here
	if len(name) < 5 {
		fmt.Println("Error! '",name,"' less than 5 chars")
	} else {
		firstChar := string(name[0])
		lastChar := string(name[len(name) - 1])
		amountChar := strconv.Itoa(len(name))

		for i := len(name) - 1; i >= 0 ; i-- {
			reverse += string(name[i])
		}

		result := strings.ToUpper(firstChar) + reverse + strings.ToUpper(lastChar) + amountChar
		fmt.Println("The password's =", result)
	}
}
