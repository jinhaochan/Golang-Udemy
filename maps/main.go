package main

import "fmt"

// a map is like a dictionary
// map keys must be of the same type
// map values must be of the same type
// map key and values need not be the same type

func main() {
	/*
		colors := map[string]string{"red": "#ff0000", "blue": "#0201021"}

		fmt.Printf("%v+\n", colors)
	*/

	// creating a map with the inbuilt function make
	colors := make(map[string]string)

	// adding items
	colors["white"] = "#000000"
	colors["red"] = "#ff0000"
	colors["blue"] = "#0201021"
	colors["green"] = "#43321"

	// deleting items
	delete(colors, "white")

	// if white exists, val will be the value, and ok = true
	// if white does not exists, val will be the zero value for string, and ok = false
	val, ok := colors["white"]

	fmt.Println(val, ok)

	printMap(colors)

}
