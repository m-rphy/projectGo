package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// receiver to shuffle the deck
func (d deck) shuffle() deck {
    // the seed of our randomness
    source := rand.NewSource(time.Now().UnixNano())
    // the random generator source
    r := rand.New(source)
    for i := range d {
        // generate new positions from r
        np := r.Intn(len(d) - 1)
        // swap positions with the card ar randNum
        d[i], d[np] = d[np], d[i]
    }
    return d
}

// receiver that turns deck -> string
func (d deck) toString() string {
    return strings.Join([]string(d), ",") 
}

// receiver to write to a file 
func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {

    bs, err := ioutil.ReadFile(filename)
    
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(2)
    }

    s := strings.Split(string(bs), ",")
    return deck(s) 
}

