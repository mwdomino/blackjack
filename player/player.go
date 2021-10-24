package player

import (
	"github.com/mwdomino/blackjack/hand"
)

// Player represents an actor in the Game with their properties
type Player struct {
	Hand hand.Hand
	Wins int
	Name string
}

// New creates a new player
func New(name string) *Player {
	return &Player{
		Hand: *hand.New(),
		Name: name,
	}
}
