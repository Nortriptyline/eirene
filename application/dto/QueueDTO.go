package dto

// import (
// 	"github.com/Nortriptyline/Eirene/domain"
// 	"github.com/google/uuid"
// )

// type QueueDTO struct {
// 	ID   uuid.UUID `json:"id"`
// 	Name string    `json:"name"`
// 	Room *RoomDTO  `json:"room"`
// }

// type QueueLightDTO struct {
// 	ID   uuid.UUID `json:"id"`
// 	Name string    `json:"name"`
// }

// func ToQueueDTO(queue *domain.Queue) *QueueDTO {
// 	return &QueueDTO{
// 		ID:   queue.ID,
// 		Name: queue.Name,
// 		Room: ToRoomDTO(queue.Room),
// 	}
// }

// func ToQueueLightDTO(queue *domain.Queue) *QueueLightDTO {
// 	return &QueueLightDTO{
// 		ID:   queue.ID,
// 		Name: queue.Name,
// 	}
// }
