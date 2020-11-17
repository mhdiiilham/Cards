package main

import (
	"fmt"
)

type anonymouseName func(string) bool

func main() {

	// 1. Function as value
	bye := sayGoodbye
	fmt.Println(bye("Vaseline"))

	// 2. Function as parameter
	filterFunction := thisOneIsAfilter
	secondResult := sayHelloWithFilter("Ilham", filterFunction)
	badWord := sayHelloWithFilter("Goblok", filterFunction)
	fmt.Println(secondResult)
	fmt.Println(badWord)

	// 3. Anonymous function
	sayHello("Muhammad Ilham", func(name string) bool {
		if name != "Muhammad Ilham" {
			return false
		}
		return true
	})
	sayHello("Zackkk", func(name string) bool {
		if name != "Muhammad Ilham" {
			return false
		}
		return true
	})

}

func sayGoodbye(name string) string {
	return "Goodbye " + name
}

func sayHelloWithFilter(name string, filter func(string) string) string {
	return "Hello " + filter(name)
}

func thisOneIsAfilter(word string) string {
	if word == "Goblok" {
		return "Orang Baik"
	}

	return word
}

func sayHello(name string, an anonymouseName) {
	if an(name) {
		fmt.Println("Halo")
	} else {
		fmt.Println("Ha?")
	}
}
