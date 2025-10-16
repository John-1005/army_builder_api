package models

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Edition   string    `json:"edition"`
	Version   string    `json:"version"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
