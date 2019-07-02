package poker

import "time"

// Player represents a player at a table
type Player struct {
	AccountID  int
	Bankroll   float64
	Nickname   string
	Avatar     string
	Connected  bool
	Inactive   bool
	LastAction time.Time
}
