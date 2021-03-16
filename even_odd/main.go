package main

import "fmt"

func main() {
	s := make([]int, 11)

	for i := range s {
		if i%2 == 0 {
			fmt.Println(i, " is Even")
		} else {
			fmt.Println(i, " is Odd")
		}
	}
}
