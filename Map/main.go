package main

import (
	"fmt"
)

func main() {
	// Declaring a map
	colors := map[string]string {
		// here's we're declaring a map
		// that all of it key are string
		// and the value is string as well
		"red": "#FF0000",
		"green": "#00FF00",
		"blue": "#0000FF",
	}
	fmt.Println(colors)
}