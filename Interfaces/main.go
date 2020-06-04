package main

import "fmt"

type bot interface{
	// Any types inside this program,
	// that have function called `getGreeting`
	// and return a string, now is an
	// honory member of bot interfaces
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}