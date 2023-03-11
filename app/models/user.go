package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// User model struct.
type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Users array model struct of User.
type Users []User

// String converts the struct into a string value.
func (u User) String() string {
	return fmt.Sprintf("%+v\n", u)
}
