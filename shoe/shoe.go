package shoe

import (
	"math/rand"
	"time"

	"github.com/mwdomino/blackjack/card"
)

// Shoe represents a grouping of many Cards for use across multiple Hands
type Shoe struct {
	Cards []card.Card
}

// addDeck creates a new 52 card deck and appends it to the Shoe
func (shoe *Shoe) addDeck() {
	var cards = make([]card.Card, 52)

	i := 0
	suits := card.AllSuits()
	for s := range suits {
		for n, v := range *card.AllValues() {
			cards[i] = card.Card{
				Name:  n,
				Value: v,
				Suit:  suits[s],
			}
			i++
		}
	}
	shoe.Cards = append(shoe.Cards, cards...)
}

// Shuffle randomizes the order of Cards in the Seck
func (shoe *Shoe) Shuffle() {
	c := shoe.Cards
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(c),
		func(i, j int) {
			c[i], c[j] = c[j], c[i]
		})
}

// DrawCard draws a single card, removes it from the Shoe and returns it
func (shoe *Shoe) DrawCard() card.Card {
	card := shoe.Cards[0]
	shoe.Cards = shoe.Cards[1:]

	return card
}

// New builds the Shoe from decks of Cards and then shuffles it
func New(decks int, shuffles int) *Shoe {
	var ret Shoe

	for i := 0; i < decks; i++ {
		ret.addDeck()
	}

	for i := 0; i < shuffles; i++ {
		ret.Shuffle()
	}
	return &ret
}
