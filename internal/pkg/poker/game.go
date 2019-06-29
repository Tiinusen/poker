package poker

import (
	"fmt"
	"time"
)

// Game represents a poker game (also tournaments)
type Game struct {
	Type  GameType
	Rules Rules
	Limit Limit

	// Break Related
	NextBreak     time.Time
	BreakInterval time.Duration

	// Mandatory Structure Related
	NextLevel          time.Time
	CurrentLevel       uint
	LevelInterval      time.Duration
	MandatoryStructure []MandatoryStructure
}

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
