package entity

import (
	"time"

	"github.com/satori/uuid"
)

type Users struct {
	Uid       uuid.UUID
	Name      string
	Email     string
	Passowrd  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
