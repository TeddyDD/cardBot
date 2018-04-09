package initiative

import (
	"fmt"
	"sort"
)

type Combat struct {
	Deck     *Deck
	Rejected *Deck
	Players  map[string]*Deck
}

func NewCombat() *Combat {
	c := new(Combat)
	c.Deck = NewDeck()
	c.Rejected = &Deck{}
	c.Players = make(map[string]*Deck)
	return c
}

func (c *Combat) NewPlayer(name string) {
	c.Players[name] = new(Deck)
}

func (c *Combat) GiveCards() {
	for _, p := range c.Players {
		if c.Deck.IsEmpty() {
			c.Reshuffle()
			fmt.Println("RESHUFFLE")
		} else {
			c.Deck.MoveRandomCard(p)
		}
	}
}

// Reshuffle cards from rejected deck
func (c *Combat) Reshuffle() {
	*c.Deck = append(*c.Deck, *c.Rejected...)
	c.Rejected = &Deck{}
	c.Deck.Shuffle()
}

type CardInfo struct {
	Name string
	Card *Card
}

type CardInfoByValue []*CardInfo

func (b CardInfoByValue) Len() int {
	return len(b)
}

func (b CardInfoByValue) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b CardInfoByValue) Less(i, j int) bool {
	if b[i].Card.value == b[j].Card.value {
		return b[i].Card.suit < b[j].Card.suit
	}
	return b[i].Card.value < b[j].Card.value
}

func (c *Combat) GetInitiative() (queue *[]string) {
	var allCards []*CardInfo
	for name, pDeck := range c.Players {
		cInfo := &CardInfo{
			Name: name,
			Card: &(*pDeck)[0],
		}
		allCards = append(allCards, cInfo)
		// remove from players deck and return to combat deck
		ret := &(*pDeck)[0]
		*pDeck = append((*pDeck)[:0], (*pDeck)[1:]...)
		*c.Rejected = append(*c.Deck, *ret)
		// shuffle combat deck just in case
		c.Deck.Shuffle()
	}
	sort.Sort(sort.Reverse(CardInfoByValue(allCards)))
	queue = new([]string)
	for _, res := range allCards {
		*queue = append(*queue, fmt.Sprintf("%s - %s", res.Name, res.Card))
	}
	return
}
