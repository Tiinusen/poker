package poker

import (
	"testing"

	"github.com/go-test/deep"
)

func TestDeck_Draw(t *testing.T) {
	tests := []struct {
		name   string
		before Deck
		after  Deck
		amount int
		want   []Card
	}{
		{
			before: Deck{
				Card{CardRank: Ace, CardSuit: Spade},
				Card{CardRank: King, CardSuit: Spade},
				Card{CardRank: Queen, CardSuit: Spade},
				Card{CardRank: Jack, CardSuit: Spade},
				Card{CardRank: Ten, CardSuit: Spade},
			},
			after: Deck{
				Card{CardRank: Queen, CardSuit: Spade},
				Card{CardRank: Jack, CardSuit: Spade},
				Card{CardRank: Ten, CardSuit: Spade},
			},
			amount: 2,
			want: []Card{
				Card{CardRank: Ace, CardSuit: Spade},
				Card{CardRank: King, CardSuit: Spade},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			after, got := tt.before.Draw(tt.amount)
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("Deck.Draw() got diff = %v", diff)
			}
			if diff := deep.Equal(after, tt.after); diff != nil {
				t.Errorf("Deck.Draw() deck diff = %v", diff)
			}
		})
	}
}
