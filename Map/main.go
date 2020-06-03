package main

import (
	"fmt"
)

func main() {
	colors := map[string]string {
		"red": "#FF0000",
		"green": "#00FF00",
		"blue": "#0000FF",
		"white": "FFFFFF",
	}
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("HEX code for", color, "is", hex)
	}
}