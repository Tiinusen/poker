package rules

import (
	"testing"

	"github.com/Tiinusen/poker/internal/pkg/poker"
	"github.com/go-test/deep"
)

func TestTexasHoldEmRules_Stage(t *testing.T) {
	type args struct {
		table poker.Table
		stage poker.Stage
	}
	tests := []struct {
		name          string
		t             TexasHoldEmRules
		args          args
		wantOutcome   poker.Table
		wantNextStage poker.Stage
	}{
		{
			name: "Deal",
			t:    TexasHoldEmRules{},
			args: args{
				stage: poker.Deal,
				table: poker.Table{
					Slots: []poker.Slot{
						poker.Slot{
							Player: poker.Player{},
							Status: poker.Playing,
						},
						poker.Slot{},
						poker.Slot{
							Player: poker.Player{},
							Status: poker.Playing,
						},
						poker.Slot{},
						poker.Slot{},
						poker.Slot{},
					},
					Deck: poker.Deck{
						poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
						poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
						poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
						poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
					},
				},
			},
			wantNextStage: poker.PreFlop,
			wantOutcome: poker.Table{
				Slots: []poker.Slot{
					poker.Slot{
						Cards: []poker.Card{
							poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
							poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
						},
						Status: poker.Playing,
					},
					poker.Slot{},
					poker.Slot{
						Cards: []poker.Card{
							poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
							poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
						},
						Status: poker.Playing,
					},
					poker.Slot{},
					poker.Slot{},
					poker.Slot{},
				},
				Deck: poker.Deck{},
			},
		},
		{
			name: "PreFlop",
			t:    TexasHoldEmRules{},
			args: args{
				stage: poker.PreFlop,
				table: poker.Table{
					Slots: []poker.Slot{
						poker.Slot{
							Cards: []poker.Card{
								poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
								poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
							},
							Status: poker.Playing,
						},
						poker.Slot{},
						poker.Slot{
							Cards: []poker.Card{
								poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
								poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
							},
							Status: poker.Playing,
						},
						poker.Slot{},
						poker.Slot{},
						poker.Slot{},
					},
					Deck: poker.Deck{
						poker.Card{CardRank: poker.Ace, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.King, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
					},
				},
			},
			wantNextStage: poker.Flop,
			wantOutcome: poker.Table{
				Slots: []poker.Slot{
					poker.Slot{
						Cards: []poker.Card{
							poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
							poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
						},
						Status: poker.Playing,
					},
					poker.Slot{},
					poker.Slot{
						Cards: []poker.Card{
							poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
							poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
						},
						Status: poker.Playing,
					},
					poker.Slot{},
					poker.Slot{},
					poker.Slot{},
				},
				DiscardedCards: []poker.Card{
					poker.Card{CardRank: poker.Ace, CardSuit: poker.Heart},
				},
				CommunityCards: []poker.Card{
					poker.Card{CardRank: poker.King, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
				},
				Deck: poker.Deck{},
			},
		},
		{
			name: "Flop",
			t:    TexasHoldEmRules{},
			args: args{
				stage: poker.Flop,
				table: poker.Table{
					Slots: []poker.Slot{
						poker.Slot{
							Cards: []poker.Card{
								poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
								poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
							},
							Status: poker.Playing,
						},
						poker.Slot{},
						poker.Slot{
							Cards: []poker.Card{
								poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
								poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
							},
							Status: poker.Playing,
						},
						poker.Slot{},
						poker.Slot{},
						poker.Slot{},
					},
					DiscardedCards: []poker.Card{
						poker.Card{CardRank: poker.Ace, CardSuit: poker.Heart},
					},
					CommunityCards: []poker.Card{
						poker.Card{CardRank: poker.King, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
					},
					Deck: poker.Deck{
						poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
						poker.Card{CardRank: poker.King, CardSuit: poker.Diamond},
						poker.Card{CardRank: poker.Queen, CardSuit: poker.Diamond},
						poker.Card{CardRank: poker.Jack, CardSuit: poker.Diamond},
					},
				},
			},
			wantNextStage: poker.Turn,
			wantOutcome: poker.Table{
				Slots: []poker.Slot{
					poker.Slot{
						Cards: []poker.Card{
							poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
							poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
						},
						Status: poker.Playing,
					},
					poker.Slot{},
					poker.Slot{
						Cards: []poker.Card{
							poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
							poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
						},
						Status: poker.Playing,
					},
					poker.Slot{},
					poker.Slot{},
					poker.Slot{},
				},
				DiscardedCards: []poker.Card{
					poker.Card{CardRank: poker.Ace, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
				},
				CommunityCards: []poker.Card{
					poker.Card{CardRank: poker.King, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.King, CardSuit: poker.Diamond},
				},
				Deck: poker.Deck{
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Diamond},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Diamond},
				},
			},
		},
		{
			name: "Turn",
			t:    TexasHoldEmRules{},
			args: args{
				stage: poker.Turn,
				table: poker.Table{
					Slots: []poker.Slot{
						poker.Slot{
							Cards: []poker.Card{
								poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
								poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
							},
							Status: poker.Playing,
						},
						poker.Slot{},
						poker.Slot{
							Cards: []poker.Card{
								poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
								poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
							},
							Status: poker.Playing,
						},
						poker.Slot{},
						poker.Slot{},
						poker.Slot{},
					},
					DiscardedCards: []poker.Card{
						poker.Card{CardRank: poker.Ace, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
					},
					CommunityCards: []poker.Card{
						poker.Card{CardRank: poker.King, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.King, CardSuit: poker.Diamond},
					},
					Deck: poker.Deck{
						poker.Card{CardRank: poker.Queen, CardSuit: poker.Diamond},
						poker.Card{CardRank: poker.Jack, CardSuit: poker.Diamond},
					},
				},
			},
			wantNextStage: poker.River,
			wantOutcome: poker.Table{
				Slots: []poker.Slot{
					poker.Slot{
						Cards: []poker.Card{
							poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
							poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
						},
						Status: poker.Playing,
					},
					poker.Slot{},
					poker.Slot{
						Cards: []poker.Card{
							poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
							poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
						},
						Status: poker.Playing,
					},
					poker.Slot{},
					poker.Slot{},
					poker.Slot{},
				},
				DiscardedCards: []poker.Card{
					poker.Card{CardRank: poker.Ace, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Diamond},
				},
				CommunityCards: []poker.Card{
					poker.Card{CardRank: poker.King, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.King, CardSuit: poker.Diamond},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Diamond},
				},
				Deck: poker.Deck{},
			},
		},
		{
			name: "River",
			t:    TexasHoldEmRules{},
			args: args{
				stage: poker.River,
				table: poker.Table{
					Slots: []poker.Slot{
						poker.Slot{
							Cards: []poker.Card{
								poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
								poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
							},
							Status: poker.Playing,
						},
						poker.Slot{},
						poker.Slot{
							Cards: []poker.Card{
								poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
								poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
							},
							Status: poker.Playing,
						},
						poker.Slot{},
						poker.Slot{},
						poker.Slot{},
					},
					DiscardedCards: []poker.Card{
						poker.Card{CardRank: poker.Ace, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
						poker.Card{CardRank: poker.Queen, CardSuit: poker.Diamond},
					},
					CommunityCards: []poker.Card{
						poker.Card{CardRank: poker.King, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
						poker.Card{CardRank: poker.King, CardSuit: poker.Diamond},
						poker.Card{CardRank: poker.Jack, CardSuit: poker.Diamond},
					},
					Deck: poker.Deck{
						poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					},
				},
			},
			wantNextStage: poker.End,
			wantOutcome: poker.Table{
				Slots: []poker.Slot{
					poker.Slot{
						Cards: []poker.Card{
							poker.Card{CardRank: poker.Ace, CardSuit: poker.Spade},
							poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
						},
						Status: poker.Playing,
					},
					poker.Slot{},
					poker.Slot{
						Cards: []poker.Card{
							poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
							poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
						},
						Status: poker.Playing,
					},
					poker.Slot{},
					poker.Slot{},
					poker.Slot{},
				},
				DiscardedCards: []poker.Card{
					poker.Card{CardRank: poker.Ace, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.Ace, CardSuit: poker.Diamond},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Diamond},
				},
				CommunityCards: []poker.Card{
					poker.Card{CardRank: poker.King, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
					poker.Card{CardRank: poker.King, CardSuit: poker.Diamond},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Diamond},
				},
				Deck: poker.Deck{
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutcome, gotNextStage := tt.t.Stage(tt.args.table, tt.args.stage)
			if diff := deep.Equal(gotOutcome, tt.wantOutcome); diff != nil {
				t.Errorf("TexasHoldEmStage() gotOutcome diff = %v", diff)
			}
			if diff := deep.Equal(gotNextStage, tt.wantNextStage); diff != nil {
				t.Errorf("TexasHoldEmStage() gotNextStage diff = %v", diff)
			}
		})
	}
}

func TestTexasHoldEmRules_Win(t *testing.T) {
	type args struct {
		player1   []poker.Card
		player2   []poker.Card
		community []poker.Card
	}
	tests := []struct {
		name string
		t    TexasHoldEmRules
		args args
		want int
	}{
		{
			name: "High Card (Tie)",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Five, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.Six, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.Eight, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Nine, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Club},
					poker.Card{CardRank: poker.King, CardSuit: poker.Club},
				},
			},
			want: 0,
		},
		{
			name: "High Card",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Ten, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.Six, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.Eight, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Nine, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.King, CardSuit: poker.Club},
				},
			},
			want: 1,
		},
		{
			name: "High Card vs Pair",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Five, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.Eight, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Nine, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.King, CardSuit: poker.Club},
				},
			},
			want: 2,
		},
		{
			name: "Two Pair vs Pair",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Eight, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.Eight, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Nine, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.King, CardSuit: poker.Club},
				},
			},
			want: 1,
		},
		{
			name: "Two Pair vs Three of a Kind",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Eight, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.Eight, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Nine, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
				},
			},
			want: 2,
		},
		{
			name: "Two Pair vs Three of a Kind",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Eight, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.Eight, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Nine, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
				},
			},
			want: 2,
		},
		{
			name: "Straight vs Three of a Kind",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Ten, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Heart},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.Ace, CardSuit: poker.Club},
					poker.Card{CardRank: poker.King, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
				},
			},
			want: 1,
		},
		{
			name: "Straight vs Flush",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Ten, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.Five, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Club},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.Ace, CardSuit: poker.Club},
					poker.Card{CardRank: poker.King, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
				},
			},
			want: 2,
		},
		{
			name: "Full House vs Flush",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.Five, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Two, CardSuit: poker.Club},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.Ace, CardSuit: poker.Club},
					poker.Card{CardRank: poker.King, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
				},
			},
			want: 1,
		},
		{
			name: "Full House vs Four of a Kind",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Heart},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.King, CardSuit: poker.Heart},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.King, CardSuit: poker.Diamond},
					poker.Card{CardRank: poker.King, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Heart},
				},
			},
			want: 2,
		},
		{
			name: "Straight Flush vs Four of a Kind",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Nine, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Eight, CardSuit: poker.Club},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.King, CardSuit: poker.Diamond},
					poker.Card{CardRank: poker.King, CardSuit: poker.Heart},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.King, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Ten, CardSuit: poker.Club},
				},
			},
			want: 1,
		},
		{
			name: "Straight Flush vs Royal Flush",
			t:    TexasHoldEmRules{},
			args: args{
				player1: []poker.Card{
					poker.Card{CardRank: poker.Nine, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Eight, CardSuit: poker.Club},
				},
				player2: []poker.Card{
					poker.Card{CardRank: poker.Ace, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Nine, CardSuit: poker.Heart},
				},
				community: []poker.Card{
					poker.Card{CardRank: poker.King, CardSuit: poker.Spade},
					poker.Card{CardRank: poker.King, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Jack, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Queen, CardSuit: poker.Club},
					poker.Card{CardRank: poker.Ten, CardSuit: poker.Club},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Win(tt.args.player1, tt.args.player2, tt.args.community); got != tt.want {
				t.Errorf("TexasHoldEmWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
