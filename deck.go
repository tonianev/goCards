package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// represents a slice of strings, a custom type for a deck of cards
type deck []string

// creates and returns a new deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// prints all the cards in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// splits the deck into two decks: one of size handSize and the other with the remaining cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// converts the deck to a single comma-separated string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saves the deck to a file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// loads a deck from a file
func newDeckFromFile(filename string) (deck, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", filename, err)
	}
	s := strings.Split(string(bs), ",")
	return deck(s), nil
}

// randomly shuffles the cards in the deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
