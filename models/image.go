package models

import (
	"time"

	"github.com/google/uuid"
)

type ImageStatus int

const (
	ImageStatusPending   ImageStatus = iota
	ImageStatusFailed    ImageStatus = iota
	ImageStatusCompleted ImageStatus = iota
)

type Image struct {
	ID        int //`bun:id,pk,autoincrement`
	UserID    uuid.UUID
	Status    ImageStatus
	CreatedAt time.Time //`bun:"default:'now()'"`

}
