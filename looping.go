package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 && i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "Fuzz", "Buzz", "FuzzBuzz")
		} else if i%2 == 0 && i%3 == 0 {
			fmt.Println(i, "Buzz", "Fuzz")
		} else if i%2 == 0 && i%5 == 0 {
			fmt.Println(i, "Buzz", "FuzzBuzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "Fuzz", "FuzzBuzz")
		} else if i%2 == 0 {
			fmt.Println(i, "Buzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fuzz")
		} else if i%5 == 0 {
			fmt.Println(i, "FuzzBuzz")
		} else {
			continue
		}
	}
}
