package main

import "fmt"

func main() {
	// your code here
	for i := 0;  i <= 100; i++ {
		
		// dapat bagi 2 show buzz
		if i % 2 == 0 {
			fmt.Println(i, ": is buzz")
			// dapat bagi 3 show fuzz
			} else if i % 3 == 0 {
				fmt.Println(i, ": is fuzz")
		// dapat bagi 5 show BuzzFuzz 
		} else if i % 5 == 0 {
			fmt.Println(i, ": is BuzzFuzz")
		}
	}
}
