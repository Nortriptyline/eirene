package application

type Queries struct {
	// GetRoomsHandler *GetRoomsQueryHandler
}

func InitializeQueries() *Queries {
	return &Queries{
		// GetRoomsHandler: NewGetRoomsQueryHandler(roomService),
	}
}
