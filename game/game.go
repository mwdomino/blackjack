package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"

	"github.com/mwdomino/blackjack/player"
	"github.com/mwdomino/blackjack/shoe"
)

// Game represents the state of the game at any point
type Game struct {
	Shoe        *shoe.Shoe
	Player      *player.Player
	Dealer      *player.Player
	LastStatus  string
	Winner      *player.Player
	HandsPlayed int
}

// New builds a new Game with Players and a Shoe
func New(decks int, shuffles int) *Game {
	return &Game{
		Shoe:   shoe.New(decks, shuffles),
		Player: player.New("Player"),
		Dealer: player.New("Dealer"),
	}
}

// DealHand begins the game by dealing cards to the players
// The first card will go to the dealer, second to the player
// and then another card to dealer and player
func (g *Game) DealHand() {
	for i := 0; i < 2; i++ {
		g.Dealer.Hand.Cards = append(g.Dealer.Hand.Cards, g.Shoe.DrawCard())
		g.Player.Hand.Cards = append(g.Player.Hand.Cards, g.Shoe.DrawCard())
	}
}

// DealCard assigns a single new card to a Player, removing it from the Shoe
func (g *Game) DealCard(p *player.Player) {
	p.Hand.Cards = append(p.Hand.Cards, g.Shoe.DrawCard())
}

// mainScreen is the screen shown during each round of play
func (g *Game) mainScreen() {
	fmt.Printf("\033c")
	fmt.Printf("\t\t\t\t\t\t\t Shoe: %d/%d\n", len(g.Shoe.Cards), 6*52)
	fmt.Printf("\t\t\t\t\t\t\t Last Move: %s\n", g.LastStatus)
	fmt.Printf("Dealer shows: %s\n", g.Dealer.Hand.Output(true))
	fmt.Printf("Player Shows (%d): %s\n", g.Player.Hand.Value(), g.Player.Hand.Output(false))
}

// finalScore shows the final screen, describing who won and with what
func (g *Game) finalScore() {
	fmt.Printf("\033c")
	fmt.Printf("After a rousing game...\n\n")
	fmt.Printf("Dealer shows: %s\n", g.Dealer.Hand.Output(false))
	fmt.Printf("Player Shows: %s\n\n", g.Player.Hand.Output(false))

	g.Winner.Wins += 1
	g.HandsPlayed += 1

	if g.IsDraw() {
		fmt.Printf("We have a DRAW! No winner this hand :(\n")
	} else {
		fmt.Printf("%s wins with a hand of %d!\n\n", g.Winner.Name, g.Winner.Hand.Value())
	}
	fmt.Printf("The running score is....\n")
	fmt.Printf("Dealer %d and Player %d after %d games!", g.Dealer.Wins, g.Player.Wins, g.HandsPlayed)

	g.ResetHands()
}

// IsDraw determines whether the game is a draw, with no winner
func (g *Game) IsDraw() bool {
	if g.Winner.Name == "Draw" {
		return true
	}
	return false
}

// ResetHands clears the cards from both players
func (g *Game) ResetHands() {
	g.Dealer.Hand.Reset()
	g.Player.Hand.Reset()
}

// Play represents a single hand gathering input until completion
func (g *Game) Play() {
	g.DealHand()
	g.LastStatus = "Cards dealt"

	g.mainScreen()
	playerComplete := false
	for playerComplete != true {
		reader := bufio.NewReader(os.Stdin)
		// Ask player what to do, looping until player stands:
		//   (h)it - take a new card
		//   (s)tand - stop drawing
		//   (q)uit - quit
		fmt.Printf("\n\n(h)it or (s)tand?: ")
		char, _, err := reader.ReadRune()
		if err != nil {
			log.Fatalf("Unable to read input: %v", err)
		}

		char = unicode.ToLower(char)
		switch char {
		// hit
		case 'h':
			g.LastStatus = "Player Hit"
			g.DealCard(g.Player)
			// if player busts, no more moves
			if g.Player.Hand.IsBusted() {
				g.LastStatus = "Player Busted"
				playerComplete = true
				g.Winner = g.Dealer
			}
		// stand
		case 's':
			g.LastStatus = "Standing"
			playerComplete = true
		// exit
		case 'q', 'x':
			g.LastStatus = "Player Quit"
			os.Exit(0)
		}
		g.mainScreen()
	}

	g.LastStatus = "Player Hand Complete"

	// Dealer does not need to play if player has already busted
	dealerComplete := g.Player.Hand.IsBusted() && g.Dealer.Hand.IsNatural()

	// After player completes, dealer plays according to his strategy
	// Dealer draws on 16 (and below)
	// Dealer stands on 17 (and above)
	for dealerComplete != true {
		if g.Dealer.Hand.Value() <= 16 {
			g.LastStatus = "Dealer Hit"
			g.DealCard(g.Dealer)
			if g.Dealer.Hand.IsBusted() {
				g.LastStatus = "Dealer Busted"
				dealerComplete = true
				g.Winner = g.Player
			}
		}

		if g.Dealer.Hand.Value() >= 17 {
			g.LastStatus = "Dealer Stand"
			dealerComplete = true
		}

		time.Sleep(500 * time.Millisecond)
	}

	// calculate winner
	d, p := g.Dealer.Hand, g.Player.Hand
	// if player bust, dealer win
	if p.IsBusted() {
		g.Winner = g.Dealer
	} else if d.IsBusted() {
		g.Winner = g.Player
	} else if d.Value() > p.Value() {
		g.Winner = g.Dealer
	} else if p.Value() > d.Value() {
		g.Winner = g.Player
	} else if p.Value() == d.Value() {
		g.Winner = player.New("Draw")
	}

	g.finalScore()
}
