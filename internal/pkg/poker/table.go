package poker

import (
	"time"

	"github.com/google/uuid"
)

// NewTable creates a table with provided capacity
func NewTable(capacity int) Table {
	uuid, _ := uuid.NewUUID()
	return Table{
		TableID: TableID(uuid.String()),
		Slots:   make([]Slot, capacity),
	}
}

// Table represents a poker table
type Table struct {
	TableID        TableID
	Slots          []Slot
	Deck           Deck
	CommunityCards []Card
	DiscardedCards []Card
}

// TableID represents an TableID
type TableID string

// Slot represents a table slot
type Slot struct {
	Player          Player
	Chips           float64
	Cards           []Card
	Status          PlayerStatus
	SittingOutSince time.Time
	BuyingInSince   time.Time
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
