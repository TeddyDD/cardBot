package main

import (
	"fmt"
	"math/rand"
)

var cardSuit = map[int]string{
	1: "♣",
	2: "♦",
	3: "♥",
	4: "♠",
}

var cardValues = map[int]string{
	1: "9",
	2: "10",
	3: "J",
	4: "Q",
	5: "K",
	6: "A",
	7: "JOKER",
}

// Card represents single card in poker deck
type Card struct {
	suit  int
	value int
}

func (c Card) String() string {
	color := []string{"♣", "♦", "♥", "♠"}
	var colorSymbol string
	// joker
	if c.value == 7 {
		if c.suit <= 2 { // 1-2
			colorSymbol = "CZERWONY"
		} else {
			colorSymbol = "CZARNY"
		}
	} else { // 3-4
		colorSymbol = fmt.Sprintf("%s", color[c.suit-1])
	}
	return fmt.Sprintf("%s %s", colorSymbol, cardValues[c.value])
}

// NewCard creates new card struct
func NewCard(value int, color int) *Card {
	c := &Card{}
	c.suit = color
	c.value = value
	return c
}

type Deck []Card

//NewDeck generates slice with Poker deck
func NewDeck() *Deck {
	d := new(Deck)
	for c := 1; c < 5; c++ {
		for v := 1; v < 7; v++ {
			*d = append(*d, *NewCard(v, c))
		}
	}
	*d = append(*d,
		*NewCard(7, 2),
		*NewCard(7, 3))
	return d
}

// CreateCards generates deck of Poker cards
func CreateCards() {
	color := []string{"♣", "♦", "♥", "♠"}
	value := []string{"9", "10", "J", "Q", "K", "A"}
	for _, c := range color {
		for _, v := range value {
			Cards = append(Cards, fmt.Sprintf("%s %s", c, v))
		}
	}
	Cards = append(Cards, "Czarny Joker")
	Cards = append(Cards, "Czerwony Joker")
}

// Sorting

// ByValue is sorted list of cards
type ByValue []Card

func (b ByValue) Len() int {
	return len(b)
}

func (b ByValue) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByValue) Less(i, j int) bool {
	if b[i].value == b[j].value {
		return b[i].suit < b[j].suit
	}
	return b[i].value < b[j].value
}

// Cards is global deck for quick draws
var Cards []string

func getRandomCard() string {
	return Cards[rand.Intn(len(Cards))]
}
