package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/w8rLoO/GO-Random"
)

type card struct {
	suit string
	value string
}

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
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen",
		"King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Deal function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert deck of cars to single strings
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// func to save deck to local storage
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Oops... Something wrong")
		fmt.Println(err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		np := randomize.Random(len(d))
		d[i], d[np] = d[np], d[i]
	}
}