package poker

import (
	"time"

	"github.com/google/uuid"
)

// NewTable creates a table with provided capacity
func NewTable(capacity int) Table {
	uuid, _ := uuid.NewUUID()
	table := Table{
		TableID: TableID(uuid.String()),
		Slots:   make([]Slot, capacity),
	}
	for i := range table.Slots {
		table.Slots[i] = NewSlot()
	}
	return table
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

// NewSlot creates a slot
func NewSlot() Slot {
	uuid, _ := uuid.NewUUID()
	return Slot{
		SlotID: SlotID(uuid.String()),
	}
}

// Slot represents a table slot
type Slot struct {
	SlotID          SlotID
	Player          Player
	Chips           float64
	Cards           []Card
	Status          PlayerStatus
	SittingOutSince time.Time
	BuyingInSince   time.Time
}

// SlotID represents an SlotID
type SlotID string

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
