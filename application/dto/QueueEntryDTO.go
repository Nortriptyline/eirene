package dto

// import (
// 	"github.com/Nortriptyline/Eirene/domain"
// )

// type QueueEntryDTO struct {
// 	ID       string         `json:"id"`
// 	User     *UserLightDTO  `json:"user"`
// 	Queue    *QueueLightDTO `json:"queue"`
// 	Position int            `json:"position"`
// 	Status   string         `json:"status"`
// }

// type QueueEntryLightDTO struct {
// 	ID       string `json:"id"`
// 	Position int    `json:"position"`
// 	Status   string `json:"status"`
// }

// func ToQueueEntryDTO(queueEntry *domain.QueueEntry) *QueueEntryDTO {
// 	return &QueueEntryDTO{
// 		ID:       queueEntry.ID.String(),
// 		User:     ToUserLightDTO(queueEntry.User),
// 		Queue:    ToQueueLightDTO(queueEntry.Queue),
// 		Position: queueEntry.Position,
// 		Status:   string(queueEntry.Status),
// 	}
// }

// func ToQueueEntryLightDTO(queueEntry *domain.QueueEntry) *QueueEntryLightDTO {
// 	return &QueueEntryLightDTO{
// 		ID:       queueEntry.ID.String(),
// 		Position: queueEntry.Position,
// 		Status:   string(queueEntry.Status),
// 	}
// }
