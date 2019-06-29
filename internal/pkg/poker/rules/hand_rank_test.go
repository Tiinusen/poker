package rules

import (
	"reflect"
	"testing"

	"github.com/Tiinusen/poker/internal/pkg/poker"
)

func TestHandRankAndTopCards(t *testing.T) {
	tests := []struct {
		name  string
		cards []poker.Card
		want  HandRank
		want1 []poker.Card
	}{
		{
			name: "Royal Flush",
			cards: []poker.Card{
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Three, CardSuit: poker.Diamond},
			},
			want: RoyalFlush,
			want1: []poker.Card{
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Spade},
			},
		},
		{
			name: "Straight Flush",
			cards: []poker.Card{
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Three, CardSuit: poker.Diamond},
			},
			want: StraightFlush,
			want1: []poker.Card{
				poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Spade},
			},
		},
		{
			name: "Four of a kind",
			cards: []poker.Card{
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Diamond},
			},
			want: FourOfAKind,
			want1: []poker.Card{
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Diamond},
			},
		},
		{
			name: "Full House",
			cards: []poker.Card{
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Diamond},
			},
			want: FullHouse,
			want1: []poker.Card{
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Diamond},
			},
		},
		{
			name: "Flush",
			cards: []poker.Card{
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Two, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Six, CardSuit: poker.Diamond},
			},
			want: Flush,
			want1: []poker.Card{
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Six, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Two, CardSuit: poker.Diamond},
			},
		},
		{
			name: "Straight",
			cards: []poker.Card{
				poker.Card{CardRank: poker.Six, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Eight, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
			},
			want: Straight,
			want1: []poker.Card{
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Eight, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Six, CardSuit: poker.Diamond},
			},
		},
		{
			name: "Straight (when Two Pair is present)",
			cards: []poker.Card{
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Club},
				poker.Card{CardRank: poker.King, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
			},
			want: Straight,
			want1: []poker.Card{
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Club},
				poker.Card{CardRank: poker.King, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Spade},
			},
		},
		{
			name: "Three of a Kind",
			cards: []poker.Card{
				poker.Card{CardRank: poker.Six, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
			},
			want: ThreeOfAKind,
			want1: []poker.Card{
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Club},
			},
		},
		{
			name: "Two Pair",
			cards: []poker.Card{
				poker.Card{CardRank: poker.Six, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Heart},
			},
			want: TwoPair,
			want1: []poker.Card{
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Heart},
			},
		},
		{
			name: "Pair",
			cards: []poker.Card{
				poker.Card{CardRank: poker.Six, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.King, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Heart},
			},
			want: Pair,
			want1: []poker.Card{
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.King, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Heart},
			},
		},
		{
			name: "High Card",
			cards: []poker.Card{
				poker.Card{CardRank: poker.Six, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.King, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Spade},
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
			},
			want: HighCard,
			want1: []poker.Card{
				poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.King, CardSuit: poker.Club},
				poker.Card{CardRank: poker.Ten, CardSuit: poker.Diamond},
				poker.Card{CardRank: poker.Nine, CardSuit: poker.Heart},
				poker.Card{CardRank: poker.Seven, CardSuit: poker.Spade},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := HandRankWithTopCards(tt.cards)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandRankAndTopCards() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HandRankAndTopCards() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
