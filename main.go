package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"

	"github.com/mwdomino/blackjack/game"
)

func main() {
	g := game.New(6, 5)
	gameEnded := false
	for gameEnded != true {
		fmt.Printf("\033c")
		fmt.Println("Starting a new game of BLACKJACK!")
		g.Play()
		fmt.Printf("\n\nWould you like to play again?\n")

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("ENTER to continue, or 'q' to quit...")
		char, _, err := reader.ReadRune()
		if err != nil {
			log.Fatalf("Unable to read input: %v", err)
		}

		char = unicode.ToLower(char)
		switch char {
		case 'n', 'x', 'q':
			gameEnded = true
		}

		if len(g.Shoe.Cards) < 10 {
			log.Fatalf("Shoe is getting empty, time for a break!")
			gameEnded = true
		}
	}
}
