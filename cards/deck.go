package main

import "fmt"

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

func deal(d deck, handSize int) (deck, deck) {
    
    return d[:handSize], d[handSize:] // == d[:3]
}


func (d deck) shuffle() deck {
    // 
    shuffledDeck := d

    return shuffledDeck
}

// "d" is a convention, but very similar to "this" or "self"
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}
