package http

import (
	"os"

	"github.com/Nortriptyline/Eirene/application"
	gql "github.com/Nortriptyline/Eirene/infrastructure/graphql"
	httpHandlers "github.com/Nortriptyline/Eirene/infrastructure/http/handlers"
	"github.com/Nortriptyline/Eirene/infrastructure/http/middleware"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func RegisterRoutes(router *gin.Engine, handlers *application.Handlers) {
	// roomsRoutes := router.Group("/rooms")
	// roomsRoutes.Use(middleware.AuthClient())

	baseRoutes := router.Group("/banks")
	baseRoutes.Use(middleware.AuthRequired())
	httpHandlers.RegisterBankRoutes(baseRoutes, handlers)

	// httpHandlers.CreateRoom(roomsRoutes, handlers.Commands.CreateRoomHandler)
	// httpHandlers.GetRooms(router, handlers.Queries.GetRoomsHandler)
	// httpHandlers.GetRoom(roomsRoutes, handlers.Queries.GetRoomHandler)

	// queueRoutes := router.Group("/queues")
	// queueRoutes.Use(middleware.AuthClient())
	// httpHandlers.CreateQueueEntry(queueRoutes, handlers.Commands.CreateQueueEntryHandler)
	// httpHandlers.CreateLeaveQueue(queueRoutes, handlers.Commands.LeaveQueueCommandHandler)
	// httpHandlers.CreateHoldQueueEntry(queueRoutes, handlers.Commands.HoldQueueEntryCommandHandler)
}

func RegisterGraphQlRoutes(router *gin.Engine) {
	schemaBytes, err := os.ReadFile("infrastructure/graphql/schema.graphql")
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(string(schemaBytes), &gql.Resolver{})
	// GraphQL handler
	router.POST("/graphql", func(c *gin.Context) {
		relayHandler := relay.Handler{Schema: schema}
		relayHandler.ServeHTTP(c.Writer, c.Request)
	})
}
