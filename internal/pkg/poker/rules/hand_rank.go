package rules

import (
	"sort"

	"github.com/Tiinusen/poker/internal/pkg/poker"
)

// HandRankWithTopCards calculates the hand rank and top cards
func HandRankWithTopCards(cards []poker.Card) (handRank HandRank, topCards []poker.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].CardRank < cards[j].CardRank
	})

	var hasColor poker.CardSuit
	var straight []poker.Card
	cardsSortedBySuit := make(map[poker.CardSuit][]poker.Card)
	cardsSortedByRank := make(map[poker.CardRank][]poker.Card)
	for _, card := range cards {
		cardsSortedBySuit[card.CardSuit] = append(cardsSortedBySuit[card.CardSuit], card)
		cardsSortedByRank[card.CardRank] = append(cardsSortedByRank[card.CardRank], card)
	}
	for suit, cards := range cardsSortedBySuit {
		if len(cards) >= 5 {
			hasColor = suit
			break
		}
	}
	if len(cards) >= 5 {
		straight = []poker.Card{cards[0]}
		for i := 1; i < len(cards); i++ {
			if cards[i].CardRank == straight[len(straight)-1].CardRank {
				continue // Ignores cards with same rank
			}
			if (straight[len(straight)-1].CardRank + 1) == cards[i].CardRank {
				straight = append(straight, cards[i])
			} else {
				straight = []poker.Card{cards[i]}
			}
			if len(straight) == 5 {
				break
			}
		}
	}
	if len(straight) != 5 {
		straight = nil
	}

	// Royal Flush
	for _, suit := range poker.CardSuits {
		topCards = []poker.Card{}
		for _, card := range cards {
			if card.CardSuit != suit {
				continue
			}
			for _, rank := range []poker.CardRank{poker.Ace, poker.King, poker.Queen, poker.Jack, poker.Ten} {
				if card.CardRank == rank {
					topCards = append(topCards, card)
					break
				}
			}
		}
		if len(topCards) == 5 {
			return RoyalFlush, topCards
		}
	}

	// Straight Flush
	for _, suit := range poker.CardSuits {
		topCards = []poker.Card{}
		for _, card := range cards {
			if card.CardSuit != suit {
				continue
			}
			topCards = append(topCards, card)
		}

		if len(topCards) >= 5 {
			vCard := topCards[0]
			c := 1
			for i := 1; i < len(topCards); i++ {
				if (vCard.CardRank + 1) == topCards[i].CardRank {
					c++
				} else {
					c = 1
				}
				vCard = topCards[i]
				if c == 5 {
					return StraightFlush, topCards[i-4 : 5]
				}
			}
		}
	}

	// Four of a Kind
	for _, cards := range cardsSortedByRank {
		if len(cards) == 4 {
			return FourOfAKind, cards
		}
	}

	var pairs [][]poker.Card
	var threeOfAKinds [][]poker.Card
	for _, cards := range cardsSortedByRank {
		switch len(cards) {
		case 2:
			pairs = append(pairs, cards)
		case 3:
			threeOfAKinds = append(threeOfAKinds, cards)
		}
	}

	switch {
	// Full House
	case len(pairs) >= 1 && len(threeOfAKinds) >= 1:
		handRank, topCards = FullHouse, append(pairs[0], threeOfAKinds[0]...)
	case hasColor > 0:
		handRank, topCards = Flush, cardsSortedBySuit[hasColor]
	case straight != nil:
		return Straight, straight
	// Three of a Kind
	case len(threeOfAKinds) >= 1:
		return ThreeOfAKind, threeOfAKinds[0]
	// Two Pair
	case len(pairs) >= 2:
		handRank, topCards = TwoPair, append(pairs[0], pairs[1]...)
	// Pair
	case len(pairs) == 1:
		handRank, topCards = Pair, pairs[0]
	// High Card
	default:
		if len(cards) < 5 {
			return HighCard, cards[:5]
		}
		return HighCard, cards[:5]
	}
	// Appends cards
	if len(topCards) < 5 {
		for _, card := range cards {
			alreadyAdded := false
			for _, topCard := range topCards {
				if topCard == card {
					alreadyAdded = true
					break
				}
			}
			if alreadyAdded {
				continue
			}
			topCards = append(topCards, card)
			if len(topCards) == 5 {
				break
			}
		}
	}
	sort.Slice(topCards, func(i, j int) bool {
		return topCards[i].CardRank < topCards[j].CardRank
	})
	return handRank, topCards
}

// HandRank represents a hand rank of an evaluated hand
type HandRank int

func (v HandRank) String() string {
	v--
	if v < 0 {
		return "Unknown"
	}
	return HandRankNames[v]
}

const (
	UnknownHandRank HandRank = iota
	RoyalFlush
	StraightFlush
	FourOfAKind
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPair
	Pair
	HighCard
)

// HandRanks is a list of available hand ranks
var HandRanks = []HandRank{
	RoyalFlush,
	StraightFlush,
	FourOfAKind,
	FullHouse,
	Flush,
	Straight,
	ThreeOfAKind,
	TwoPair,
	Pair,
	HighCard,
}

// HandRankNames is a list of hand rank names
var HandRankNames = []string{
	"Royal flush",
	"Straight flush",
	"Four of a kind",
	"Full house",
	"Flush",
	"Straight",
	"Three of a kind",
	"Two pair",
	"Pair",
	"High Card",
}
