package queryhandlers

type GetRoomQueryHandler struct {
	// roomService *service.RoomService
}

// func NewGetRoomQueryHandler(service *service.RoomService) *GetRoomQueryHandler {
// 	return &GetRoomQueryHandler{roomService: service}
// }

// func (h *GetRoomQueryHandler) Handle(q queries.GetRoomQuery) (*dto.RoomDTO, error) {
// 	room, err := h.roomService.GetRoom(q.ID)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return dto.ToRoomDTO(room), nil
// }
