package main

import "fmt"

func main() {

    cards :=  newDeck()
    
    // returns two values of type deck
    hand, remainingDeck := deal(cards, 5)
    fmt.Println(hand, remainingDeck)
}
