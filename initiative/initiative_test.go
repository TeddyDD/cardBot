package initiative

import (
	"testing"
)

func TestEmptyCombat(t *testing.T) {
	c := NewCombat()
	if len(*c.Deck) == 0 {
		t.Error("Empty global deck in combat")
	}
	if len(*c.Rejected) != 0 {
		t.Error("Rejected should be empty")
	}
}

func TestSinglePlayer(t *testing.T) {
	c := NewCombat()
	c.NewPlayer("a")
	if len(c.Players) != 1 {
		t.Error("Wrong ammount of players")
	}
	if _, ok := c.Players["a"]; !ok {
		t.Error("Player not added to Players map")
	}
	if len(*c.Players["a"]) != 0 {
		t.Error("Player should not have any cards")
	}
	// give cards
	c.GiveCards()
	if len(*c.Players["a"]) != 1 {
		t.Error("Player should have one card")
	}
	// initiative
	ini := c.GetInitiative()
	if len(*ini) != 1 {
		t.Error("Get Initiative should return one result")
	}
	if len(*c.Rejected) != 1 || len(*c.Players["a"]) != 0 {
		t.Error("Card should be rejected after initiative")
	}

}
