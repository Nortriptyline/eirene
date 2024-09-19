package commandhandlers

type CreateQueueEntryCommandHandler struct {
}

// func NewCreateQueueEntryCommandHandler(service *service.QueueService) *CreateQueueEntryCommandHandler {
// 	return &CreateQueueEntryCommandHandler{service: service}
// }

// func (h *CreateQueueEntryCommandHandler) Handle(cmd commands.CreateQueueEntryCommand) (*dto.QueueEntryDTO, error) {
// 	entry, err := h.service.RegisterQueueEntry(cmd.ClientToken, cmd.QueueID)

// 	// Let the caller handle the error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return dto.ToQueueEntryDTO(entry), nil
// }
