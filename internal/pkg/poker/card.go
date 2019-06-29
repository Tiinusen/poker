package poker

import "fmt"

// Card represents a playing card in a poker hand
type Card struct {
	CardRank CardRank
	CardSuit CardSuit
}

// CardSuit represents the card suit such as heart, dimaond, club and spade
type CardSuit int

func (v CardSuit) String() string {
	for i, rank := range CardSuits {
		if rank == v {
			return CardSuitNames[i]
		}
	}
	return "Unknown"
}

const (
	UnknownCardSuit CardSuit = iota
	Heart
	Diamond
	Club
	Spade
)

// CardSuits is a list of available suits
var CardSuits = []CardSuit{
	Heart,
	Diamond,
	Club,
	Spade,
}

// CardSuitNames is a list of suit names
var CardSuitNames = []string{
	"Heart",
	"Diamond",
	"Club",
	"Spade",
}

// CardRank represents the card rank such as (one, two, three, four, five, six, seven, eight, nine, ten, jack, queen, king, ace, joker)
type CardRank int

func (v CardRank) String() string {
	for i, rank := range CardRanks {
		if rank == v {
			return CardRankNames[i]
		}
	}
	return "Unknown"
}

const (
	UnknownCardRank CardRank = iota
	Ace
	King
	Queen
	Jack
	Ten
	Nine
	Eight
	Seven
	Six
	Five
	Four
	Three
	Two
	One
)

// CardRanks is a list of available suits
var CardRanks = []CardRank{
	One,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
	Ace,
}

// CardRankNames is a list of suit names
var CardRankNames = []string{
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Jack",
	"Queen",
	"King",
	"Ace",
}

// Cards can represent a hand or a deck
type Cards []Card

// Extract extracts one card out of the card collection
func (v Cards) Extract(index int) (Cards, Card, error) {
	if index >= len(v) {
		return nil, Card{}, fmt.Errorf("Index out of card collection bound")
	}
	card := v[index]
	return append(v[:index], v[index+1:]...), card, nil
}
