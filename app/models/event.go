package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

type State int

const (
	StatusConfirmed State = iota
	StatusCanceled
	StatusPending
)

func (s State) String() string {
	switch s {
	case StatusConfirmed:
		return "confirmed"
	case StatusCanceled:
		return "canceled"
	case StatusPending:
		return "pending"
	default:
		return strconv.Itoa(int(s))
	}
}

// Event model struct.
type Event struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Start     time.Time `json:"start" db:"start"`
	End       time.Time `json:"end" db:"end"`
	Status    int       `json:"status" db:"status"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	EventID   uuid.UUID `json:"event_id" db:"event_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Events array model struct of Event.
type Events []Event

// String converts the struct into a string value.
func (e Event) String() string {
	return fmt.Sprintf("%+v\n", e)
}

func (e *Event) Validate() *validate.Errors {
	return validate.Validate([]validate.Validator{
		&validators.StringIsPresent{Field: e.Name, Message: "Name cannot be blank"},
		&validators.TimeIsPresent{Field: e.Start, Message: "Start cannot be blank"},
		&validators.TimeIsPresent{Field: e.Start, Message: "Start cannot be blank"},
	}...)
}
