package main

import "fmt"

func main() {
	// your code here
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("Buzz", i)
		}
		if i%3 == 0 {
			fmt.Println("Fuzz", i)
		}
		if i%5 == 0 {
			fmt.Println("BuzzFuzz", i)
		}
	}
}
