package main

import (
	"sort"
)

import "fmt"

// Player keeps data required to initiatives tests in
// Savage Worlds

type Combat struct {
	Deck     *Deck
	Rejected *Deck
	Players  map[string]*Deck
}

func NewCombat() *Combat {
	c := new(Combat)
	c.Deck = NewDeck()
	c.Players = make(map[string]*Deck)
	return c
}

func (c *Combat) NewPlayer(name string) {
	c.Players[name] = new(Deck)
}

func (c *Combat) GiveCards() {
	for name, p := range c.Players {
		c.Deck.MoveRandomCard(p)
		if p.IsEmpty() {
			fmt.Printf("%s wtf?", name)
		} else {
			fmt.Printf("%s ma %s", name, p)
			// TODO reshuffle
		}
	}
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
		//for _, card := range *pDeck {
		//ci := &CardInfo{
		//Name: name,
		//Card: &card,
		//}
		//allCards = append(allCards, ci)
		//}
		cInfo := &CardInfo{
			Name: name,
			Card: &(*pDeck)[0],
		}
		allCards = append(allCards, cInfo)
		// remove from players deck
		*pDeck = append((*pDeck)[:0], (*pDeck)[1:]...)
	}
	sort.Sort(sort.Reverse(CardInfoByValue(allCards)))
	queue = new([]string)
	for _, res := range allCards {
		*queue = append(*queue, fmt.Sprintf("%s - %s", res.Name, res.Card))
	}
	return
}
