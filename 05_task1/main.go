package main

import "fmt"

func main() {
	// init
	s := []int{}
	for i := 0; i <= 10; i++ {
		s = append(s, i)
	}

	/// print even and odd
	for _, value := range s {
		if value%2 == 0 {
			fmt.Println(value, "is", "even")
		} else {
			fmt.Println(value, "is", "odd")
		}
	}
}
