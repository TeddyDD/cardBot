package initiative

import (
	"testing"
)

func TestMoveCard(t *testing.T) {
	expected := &Deck{}
	source := NewDeck()
	start_l := len(*source)
	err := source.MoveRandomCard(expected)
	if err != nil {
		t.Error(err)
	}
	if len(*expected) != 1 {
		t.Error("Card not moved")
	}
	l := len(*source)
	if l != start_l-1 {
		t.Errorf("Card not removed from source deck, len got %d", l)
	}
}

func TestIsEmpty(t *testing.T) {
	d := &Deck{}
	l := len(*d)
	if l != 0 || !d.IsEmpty() {
		t.Error("Deck should be empty")
	}
	if err := d.MoveRandomCard(&Deck{}); err == nil {
		t.Error("MoveRandomCard should return error on empty deck")
	}
	d = NewDeck()
	l = len(*d)
	if l == 0 || d.IsEmpty() {
		t.Error("Deck should not be empty")
	}
}

func TestShuffle(t *testing.T) {
	d := NewDeck()
	l := len(*d)
	d.Shuffle()
	if len(*d) != l {
		t.Error("Lost/added some cards in shuffle")
	}
}

func TestCardString(t *testing.T) {
	joker := NewCard(15, 3)
	if joker.String() != "CZARNY JOKER" {
		t.Errorf("Wrong joker desc, got %s", joker.String())
	}
	joker = NewCard(15, 2)
	if joker.String() != "CZERWONY JOKER" {
		t.Errorf("Wrong joker desc, got %s", joker.String())
	}
	card := NewCard(2, 1)
	if desc := card.String(); desc != "♣ 2" {
		t.Errorf("Wrong card, got %s", desc)
	}

}
