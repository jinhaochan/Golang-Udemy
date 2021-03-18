package main

import "fmt"

func printMap(c map[string]string) {
	// when iterating over a map, the two values are the key and value
	for color, hex := range c {
		fmt.Println("Hexcode for", color, "is", hex)
	}
}
