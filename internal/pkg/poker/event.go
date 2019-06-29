package poker

type Event struct {
	TableEvent *TableEvent
}

// TableEvent describes a state change on a table
type TableEvent struct {
	RoomID int
	Text   string
}

// TableEventType type of table event
type TableEventType int

const (
	UnknownTableEventType TableEventType = iota
	DeniedEntrance
)
