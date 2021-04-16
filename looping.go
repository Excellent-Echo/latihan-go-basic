package main

import "fmt"

func main() {
	// your code here
	for numberLoop := 1; numberLoop <= 100; numberLoop++ {
		if numberLoop % 2 == 0 {
			fmt.Println("Buzz")
		} else if numberLoop % 3 == 0 {
			fmt.Println("Fuzz")
		} else if numberLoop % 5 == 0 {
			fmt.Println("BuzzFuzz")
		} else {
			fmt.Println(numberLoop)
		}
	}
}
