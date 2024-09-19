package dto

// import (
// 	"github.com/Nortriptyline/Eirene/domain"
// 	"github.com/google/uuid"
// )

// type RoomDTO struct {
// 	ID       uuid.UUID        `json:"id"`
// 	Name     string           `json:"name"`
// 	Capacity int              `json:"capacity"`
// 	Queues   []*QueueLightDTO `json:"queues"`
// }

// type RoomLightDTO struct {
// 	ID       uuid.UUID `json:"id"`
// 	Name     string    `json:"name"`
// 	Capacity int       `json:"capacity"`
// }

// func ToRoomDTO(room *domain.Room) *RoomDTO {
// 	queues := make([]*QueueLightDTO, 0)
// 	for _, queue := range room.Queues {
// 		queues = append(queues, ToQueueLightDTO(queue))
// 	}

// 	return &RoomDTO{
// 		ID:       room.ID,
// 		Name:     room.Name,
// 		Capacity: room.Capacity,
// 		Queues:   queues,
// 	}
// }

// func ToRoomLightDTO(room *domain.Room) *RoomLightDTO {
// 	return &RoomLightDTO{
// 		ID:       room.ID,
// 		Name:     room.Name,
// 		Capacity: room.Capacity,
// 	}
// }
