package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string
	colors := make(map[string]string)

	// colors := map[string]string {
	// 	"red": "#FF0000",
	// 	"green": "#00FF00",
	// 	"blue": "#0000FF",
	// }

	key := "white"
	colors[key] = "#FFFFFF"	
	colors["black"] = "#000000"
	fmt.Println(colors)
	delete(colors, "black")
	fmt.Println(colors)
}