package main

import (
	"io/ioutil"
	"strings"
)

// Initializinga type "deck"
// this will be the basic type
// for the deck of cards
type deck []string

func newDeck() deck {
    cards := deck{}
    cardSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
    cardValues := []string{"Ace", "Two", "Three",
                            "Four", "Five", "Six",
                            "Seven", "Eight", "Nine",
                            "Ten", "Jack", "Queen", "King"}

    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value + " of " + suit)
        }
    }
    return cards
}

// receiver that turns deck -> string
func (d deck) toString() string {
    return strings.Join([]string(d), ",") 
}

// receiver to write to a file 
func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}



