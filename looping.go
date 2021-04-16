package main

import "fmt"

func main() {
	// your code here
	for i := 1; i <= 100; i++ {
		if i % 5 == 0 {
			fmt.Println(i, " = BuzzFuzz")
		} else if i % 3 == 0 {
			fmt.Println(i, " = Fuzz")
		} else if i % 2 == 0 {
			fmt.Println(i, " = Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
