package main

import "fmt"

// Create a new type of 'deck'
// Which is a slice of strings.
type deck []string

// A function than will
// loop through deck of cards
// and print out each on of that card.
func (d deck) print() {
	// `d` is the actual copy of the deck (custom type).
	// `deck`, every variable of type `deck` can call thid function on itself.
	for i, card := range d {
    fmt.Println(i+1, card)
  }
}

// Create and return a list of playing card
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Ace", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}
