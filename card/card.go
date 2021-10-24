package card

import "fmt"

// Card represents a single Card with a Suit, Value, and Name (text on card face)
type Card struct {
	Name  string
	Value int
	Suit  string
}

// AllSuits returns the 4 suits used in a standard deck
func AllSuits() *[4]string {
	return &[4]string{"♣", "♦", "♥", "♠"}
}

// AllValues returns all values used in a standard deck
func AllValues() *map[string]int {
	return &map[string]int{
		"A":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
		"J":  10,
		"Q":  10,
		"K":  10,
	}
}

// Stringify represents a single Card as a string of <Name><Suit>
func (c *Card) Stringify() string {
	return fmt.Sprintf("%s%s ", c.Name, c.Suit)
}
