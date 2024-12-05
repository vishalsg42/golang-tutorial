package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func deal(d deck, handSize uint) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// func saveFile(content string, filename string) {
// 	// fmt.Println([]byte(greeting))
// 	// for _, card := range d {
// 	// 	fmt.Println([]byte(card))
// 	// }
// }

// func (d deck)

// func toString(card string) byte {
// 	return []byte(card)
// }

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) saveFile(filename string) {
	os.WriteFile(filename, d.toByteSlice(), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	parsedDeck := strings.Split(string(bs), "\n")
	return deck(parsedDeck)
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		// println("New Position: ", newPosition)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
	return d
}
