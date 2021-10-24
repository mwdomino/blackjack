package hand

import (
	"fmt"

	"github.com/mwdomino/blackjack/card"
)

// Hand represents the multiple Cards in play for a single Player
type Hand struct {
	Cards []card.Card
}

// New creates and returns a new *Hand
func New() *Hand {
	return &Hand{
		Cards: make([]card.Card, 0),
	}
}

// Output returns a human readable string representing the cards
// the hideFirst argument determines if the first card is shown
//   (used for the Dealer's card being hidden during play)
func (h *Hand) Output(hideFirst bool) string {
	var ret string
	for i, c := range h.Cards {
		if i == 0 && hideFirst {
			ret += "🂠 "
			continue
		}
		ret += fmt.Sprintf("%s ", c.Stringify())
	}

	return ret
}

// aceInHand returns a boolean describing whether or not an Ace card
// is present in the given Hand
func (h *Hand) aceInHand() bool {
	for _, c := range h.Cards {
		if c.Name == "A" {
			return true
		}
	}

	return false
}

// Value returns the cumulative value of the cards in the Hand
//   if an Ace is present it returns the highest value (soft or hard) that is not
//   a bust
func (h *Hand) Value() int {
	var v int
	for _, c := range h.Cards {
		v += c.Value
	}

	// use the highest value if not a bust and an ace is in hand
	if h.aceInHand() {
		if v+10 <= 21 {
			v += 10
		}
	}

	return v
}

// Reset clears cards from the Hand
func (h *Hand) Reset() {
	h.Cards = nil
}

// IsNatural returns a boolean representing whether we have a "natural" win
//   this takes place when a player receives 21 off the deal
func (h *Hand) IsNatural() bool {
	return h.Value() == 21
}

// IsBusted returns whether or not the hand is over 21, and a bust
func (h *Hand) IsBusted() bool {
	return h.Value() > 21
}
