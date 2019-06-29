package rules

import (
	"github.com/Tiinusen/poker/internal/pkg/poker"
)

// TexasHoldEm Rules
var TexasHoldEm poker.Rules = &TexasHoldEmRules{}

// TexasHoldEmRules ...
type TexasHoldEmRules struct{}

// Stage perfoms the provided stage, adjusts accordingly and then returns the changed table state and the next stage
func (t TexasHoldEmRules) Stage(table poker.Table, stage poker.Stage) (outcome poker.Table, nextStage poker.Stage) {
	var cards []poker.Card
	switch stage {
	case poker.Deal:
		for i := 1; i <= 2; i++ {
			for i, slot := range table.Slots {
				if slot.Status != poker.Playing {
					continue
				}
				table.Deck, cards = table.Deck.Draw(1)
				table.Slots[i].Cards = append(table.Slots[i].Cards, cards...)
			}
		}
		return table, poker.PreFlop
	case poker.PreFlop:
		table.Deck, cards = table.Deck.Draw(1)
		table.DiscardedCards = append(table.DiscardedCards, cards...)
		table.Deck, cards = table.Deck.Draw(3)
		table.CommunityCards = append(table.CommunityCards, cards...)
		return table, poker.Flop
	case poker.Flop:
		table.Deck, cards = table.Deck.Draw(1)
		table.DiscardedCards = append(table.DiscardedCards, cards...)
		table.Deck, cards = table.Deck.Draw(1)
		table.CommunityCards = append(table.CommunityCards, cards...)
		return table, poker.Turn
	case poker.Turn:
		table.Deck, cards = table.Deck.Draw(1)
		table.DiscardedCards = append(table.DiscardedCards, cards...)
		table.Deck, cards = table.Deck.Draw(1)
		table.CommunityCards = append(table.CommunityCards, cards...)
		return table, poker.River
	case poker.River:
		break
	}
	return table, poker.End
}

// Win checks if player1 beats player2 with provided community cards
// returns 0 if it's a tie, 1 if player1 wins, 2 if player2 wins
func (t TexasHoldEmRules) Win(player1 []poker.Card, player2 []poker.Card, community []poker.Card) int {
	player1Rank, player1Cards := HandRankWithTopCards(append(player1, community...))
	player2Rank, player2Cards := HandRankWithTopCards(append(player2, community...))
	switch {
	case player1Rank < player2Rank:
		return 1
	case player2Rank < player1Rank:
		return 2
	default:
		// Compares top cards since both players has the same rank
		for i := 0; i < len(player1Cards); i++ {
			if player1Cards[i].CardRank == player2Cards[i].CardRank {
				continue
			}
			if player1Cards[i].CardRank > player2Cards[i].CardRank {
				return 1
			}
			return 2
		}
		return 0
	}
}
