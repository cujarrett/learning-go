package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

// "Constructor" to build a new deck of playing cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// print function that prints all cards in the deck
// It has a receiver d of type deck
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// deal function that splits the deck into two parts based on the way it's called
// It accepts two arguments, the first being d of type deck and the second handleSize of type int
// It has two returns, each of type deck
func deal(d deck, handleSize int) (deck, deck) {
	return d[:handleSize], d[handleSize:]
}

// toString function to print each card in the deck seperated by a comma and a space
// It has a receiver d of type deck
// It returns a string type
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

// saveToFile function to write the deck of cards to a file using ioutil's string byte of the string
// It has a receiver d of type deck
// It accepts a single argument called filename of type string
// It returns a error type (which comes from ioutil's WriteFile)
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, errorObject := ioutil.ReadFile(filename)
	if errorObject != nil {
		fmt.Println("Error:", errorObject)
		os.Exit(1)
	}

	sliceOfStrings := strings.Split(string(byteSlice), ", ")
	return deck(sliceOfStrings)
}
