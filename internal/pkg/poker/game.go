package poker

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// NewCashGame creates a new cash game with provided settings
func NewCashGame(Rules Rules, Limit Limit, mandatoryStructure MandatoryStructure) *Game {
	return &Game{
		Type:  Cash,
		Rules: Rules,
		Limit: Limit,
		MandatoryStructure: []MandatoryStructure{
			mandatoryStructure,
		},
	}
}

// NewTournament creates a new tournament with provided settings
func NewTournament(Rules Rules, Limit Limit, MandatoryStructure []MandatoryStructure, LevelInterval time.Duration, BreakInterval time.Duration, StartTime time.Time) *Game {
	return &Game{
		Type:               Tournament,
		Rules:              Rules,
		Limit:              Limit,
		MandatoryStructure: MandatoryStructure,
		LevelInterval:      LevelInterval,
		BreakInterval:      BreakInterval,
		StartTime:          StartTime,
	}
}

// Game represents a poker game (also tournaments)
type Game struct {
	Type  GameType
	Rules Rules
	Limit Limit

	Tables             []Table
	Players            []Player
	MandatoryStructure []MandatoryStructure

	// Break related
	NextBreak     time.Time
	BreakInterval time.Duration

	// Tournament related
	StartTime     time.Time
	CurrentLevel  uint
	NextLevel     time.Time
	LevelInterval time.Duration

	// Internal
	mutex sync.RWMutex
}

// Run should run the game until it ends or until context expires (since cash games never expires)
func (g *Game) Run(ctx context.Context) {
	// Don't forget to perform a gracevul pause of an ongoing tournament if context expires
	switch g.Type {
	case Cash:

	case Tournament:

	}
}

// Register adds player to tournament roster
// will throw ErrNotAnTournament if not an tournament
func (g *Game) Register(Player Player) error {
	if g.Type != Tournament {
		return ErrNotAnTournament
	}
	g.mutex.Lock()
	g.Players = append(g.Players, Player)
	g.mutex.Unlock()
	return nil
}

// ErrNotAnTournament is returned when trying to register to an non tournament game
var ErrNotAnTournament = errors.New("this is not a tournament")

// Join adds player to a table on designated slot
func (g *Game) Join(ctx context.Context, tableID TableID, slot int, player Player) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	for i := range g.Tables {
		if g.Tables[i].TableID == tableID {
			if slot < 0 || slot >= len(g.Tables[i].Slots) {
				return ErrTableNotFound
			}

			// Checks if Player already exists on table
			for _, p := range g.Tables[i].Slots {
				if p.Player.AccountID == player.AccountID {
					return ErrPlayerAlreadyOnTable
				}
			}

			if g.Tables[i].Slots[slot].Status != Empty {
				return ErrTableSlotAlreadyTaken
			}

			g.Tables[i].Slots[slot].Player = player
			g.Tables[i].Slots[slot].Chips = 0
			g.Tables[i].Slots[slot].Status = BuyIn
			g.Tables[i].Slots[slot].BuyingInSince = time.Now()
			return nil
		}
	}
	return ErrTableNotFound
}

// ErrPlayerAlreadyOnTable is returned when a player already plays on the same table
var ErrPlayerAlreadyOnTable = errors.New("player already added to table")

// ErrTableSlotAlreadyTaken is returned when a join attempt is made on an non empty table slot
var ErrTableSlotAlreadyTaken = errors.New("table slot is already taken")

// ErrTableNotFound is returned when no table was found
var ErrTableNotFound = errors.New("table was not found")

// Rules implements the funcs used by Game to determine outcome
type Rules interface {
	// Stage perfoms the provided stage, adjusts accordingly and then returns the changed table state and the next stage
	Stage(table Table, stage Stage) (outcome Table, nextStage Stage)
	// Win checks if slot1 beats slot2 with provided community cards
	// returns 0 if it's a tie, 1 if slot1 wins, 2 if slot2 wins
	Win(player1 []Card, player2 []Card, community []Card) int
}

// Stage represents the stage a round is in
type Stage int

const (
	ThrowTwoCard Stage = -2
	ThrowOneCard Stage = -1
	Deal         Stage = iota
	PreFlop
	Flop
	Turn
	River
	End
)

// MandatoryStructure ...
type MandatoryStructure struct {
	Blind Blind
	Ante  float64
}

// Blind represents both blinds (since they are directly related)
type Blind float64

func (v Blind) String() string {
	return fmt.Sprintf("%f/%f", v.Small(), v.Big())
}

// Big gets big blind
func (v Blind) Big() float64 {
	return float64(v)
}

// Small gets small blind
func (v Blind) Small() float64 {
	return float64(v) / 2
}

// GameType represents game type
type GameType int

func (v GameType) String() string {
	return GameTypeNames[v]
}

const (
	UnknownType GameType = iota
	Cash
	Tournament
)

// GameTypes is a list of all available game types
var GameTypes = []GameType{
	Cash,
	Tournament,
}

// LimitNames is a list of limit names
var GameTypeNames = []string{
	"Cash",
	"Tournament",
}

// Limit represents betting limit
type Limit int

func (v Limit) String() string {
	return LimitNames[v]
}

const (
	NoLimit Limit = iota
	PotLimit
	FixedLimit
)

// Limits is a list of all available limits
var Limits = []Limit{
	NoLimit,
	PotLimit,
	FixedLimit,
}

// LimitNames is a list of limit names
var LimitNames = []string{
	"No limit",
	"Pot limit",
	"Fixed limit",
}

// Blinds represents both small and big blind
type Blinds struct {
	Small float64
	Big   float64
}
