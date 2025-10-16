package models

type ReactionChoice string

const (
	ChoiceGoing    ReactionChoice = "going"
	ChoiceNotGoing ReactionChoice = "not_going"
	ChoiceThinking ReactionChoice = "thenking"
)

type Reaction struct {
	UserID    int64          `db:"user_id"`
	EventID   string         `db:"event_id"`
	Choice    ReactionChoice `db:"choice"`
	UpdatedAt string         `db:"updated_at"`
}
