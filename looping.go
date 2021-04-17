package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("FIZZ \n")
		}

		if i%3 == 0 {
			fmt.Printf("BUZZ \n")
		}

		if i%5 == 0 {
			fmt.Printf("BUZZ FUZZ \n")
		}
	}
}
