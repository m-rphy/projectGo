package main

import (
    "os"
    "testing"
)


func TestNewDeck(t *testing.T) {
    d := newDeck()
    
    // make sure length is correct
    if len(d) != 52 {
        t.Errorf("Expected length of 52, but got %v", len(d))
    }
}

// write file to disk
// (we need to clean up after we're done)
func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
    os.Remove("_deckTesting.txt")

    d := newDeck()
    d.saveToFile("_deckTesting.txt")

    ld := newDeckFromFile("_deckTesting.txt")

    if len(ld) != 52 {
        t.Errorf("Expected length of 52, got %v", len(ld))
    }

    os.Remove("_deckTesting.txt")
}
