package poker

import "time"

// NewTable creates a table with provided capacity
func NewTable(capacity int) Table {
	return Table{
		Slots: make([]Slot, capacity),
	}
}

// Table represents a poker table
type Table struct {
	Slots          []Slot
	Deck           Deck
	CommunityCards []Card
	DiscardedCards []Card
}

// Slot represents a table slot
type Slot struct {
	Player          Player
	Cards           []Card
	Status          PlayerStatus
	SittingOutSince time.Time
}

// PlayerStatus represents the players current table status
type PlayerStatus int

const (
	EnteringNextHand     PlayerStatus = -3
	EnteringNextBigBlind PlayerStatus = -2
	SittingOut           PlayerStatus = -1
	Empty                PlayerStatus = iota
	BuyIn
	Playing
	Folded
)
