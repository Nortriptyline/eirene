package queries

import "github.com/google/uuid"

type GetRoomQuery struct {
	ID uuid.UUID `json:"id"`
}
