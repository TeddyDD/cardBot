package initiative

import (
	"errors"
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
	11: "J",
	12: "Q",
	13: "K",
	14: "A",
	15: "JOKER",
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
	if c.value == 15 {
		if c.suit <= 2 { // 1-2
			colorSymbol = "CZERWONY"
		} else {
			colorSymbol = "CZARNY"
		}
	} else { // 3-4
		colorSymbol = fmt.Sprintf("%s", color[c.suit-1])
	}
	if c.value < 11 {
		return fmt.Sprintf("%s %d", colorSymbol, c.value)
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

//NewDeck generates slice with full deck of cards
func NewDeck() *Deck {
	d := new(Deck)
	for c := 1; c < 5; c++ {
		for v := 2; v < 15; v++ {
			*d = append(*d, *NewCard(v, c))
		}
	}
	*d = append(*d,
		*NewCard(15, 2),
		*NewCard(15, 3))
	return d
}

func (d Deck) IsEmpty() bool {
	return len(d) == 0
}

// MoveRandomCard from one deck to another
func (d *Deck) MoveRandomCard(to *Deck) error {
	//cardId := rand.Intn(len(*d))
	l := len(*d)
	var cardId int
	if l == 0 {
		return errors.New("Deck is empty")
	}

	cardId = rand.Intn(l)
	return d.MoveCard(cardId, to)
}

// Move card with given index to other deck
func (d *Deck) MoveCard(idx int, to *Deck) error {
	if len(*d) == 0 {
		return errors.New("Deck is empty")
	}
	if !(idx >= 0 && idx < len(*d)) {
		return errors.New("Index out of bounds")
	}
	*to = append(*to, (*d)[idx])
	*d = append((*d)[:idx], (*d)[idx+1:]...)
	return nil
}

// Shuffle deck
func (d *Deck) Shuffle() {
	for i := range *d {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}
