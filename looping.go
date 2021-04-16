package main

import "fmt"

func main() {
	// your code here
	// di bagi 2 = buz
	// di bagi 3 = fuzz
	// di bagi 5 = buzzfuzz
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("Buzz")
		}
		if i%3 == 0 {
			fmt.Println("Fuzz")
		}
		if i%5 == 0 {
			fmt.Println("BuzzFuzz")
		}
	}
}
