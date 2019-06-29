package poker

import (
	"reflect"
	"testing"
)

func TestCards_Extract(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name      string
		cards     Cards
		index     int
		wantCards Cards
		wantCard  Card
		wantErr   bool
	}{
		{
			cards: Cards{
				Card{CardRank: Ace, CardSuit: Spade},
				Card{CardRank: King, CardSuit: Spade},
				Card{CardRank: Queen, CardSuit: Spade},
				Card{CardRank: Jack, CardSuit: Spade},
				Card{CardRank: Ten, CardSuit: Spade},
			},
			index:    3,
			wantCard: Card{CardRank: Jack, CardSuit: Spade},
			wantCards: Cards{
				Card{CardRank: Ace, CardSuit: Spade},
				Card{CardRank: King, CardSuit: Spade},
				Card{CardRank: Queen, CardSuit: Spade},
				Card{CardRank: Ten, CardSuit: Spade},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.cards.Extract(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cards.Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantCards) {
				t.Errorf("Cards.Extract() got = %v, want %v", got, tt.wantCards)
			}
			if !reflect.DeepEqual(got1, tt.wantCard) {
				t.Errorf("Cards.Extract() got1 = %v, want %v", got1, tt.wantCard)
			}
		})
	}
}
