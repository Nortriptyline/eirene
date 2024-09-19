package handlers

// import (
// 	"github.com/Nortriptyline/Eirene/application/commands"
// 	"github.com/Nortriptyline/Eirene/application/handlers/commandhandlers"
// 	"github.com/Nortriptyline/Eirene/application/handlers/queryhandlers"
// 	"github.com/Nortriptyline/Eirene/application/queries"
// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// func CreateRoom(router *gin.RouterGroup, commandHandler *commandhandlers.CreateRoomCommandHandler) {
// 	router.POST("", func(c *gin.Context) {
// 		// Transform HTTP request to command
// 		command := commands.CreateRoomCommand{}
// 		if err := c.BindJSON(&command); err != nil {
// 			// TODO: log error

// 			c.JSON(400, gin.H{"error": err.Error()})
// 			return
// 		}

// 		room, err := commandHandler.Handle(command)

// 		if err != nil {
// 			c.JSON(500, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(201, gin.H{"room": room})
// 	})
// }

// func GetRooms(router *gin.Engine, commandHandler *queryhandlers.GetRoomsQueryHandler) {
// 	router.GET("/rooms", func(c *gin.Context) {
// 		// Transform HTTP request to command
// 		// command := command.GetRoomsCommand{}

// 		rooms, err := commandHandler.Handle()

// 		if err != nil {
// 			c.JSON(500, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(200, gin.H{"rooms": rooms})
// 	})
// }

// func GetRoom(router *gin.RouterGroup, commandHandler *queryhandlers.GetRoomQueryHandler) {
// 	router.GET("/:id", func(c *gin.Context) {
// 		// Transform HTTP request to command
// 		id, err := uuid.Parse(c.Param("id"))
// 		if err != nil {
// 			// Log unable to parse UUID
// 			c.JSON(400, gin.H{"error": err.Error()})
// 			return
// 		}

// 		command := queries.GetRoomQuery{
// 			ID: id,
// 		}

// 		room, err := commandHandler.Handle(command)

// 		if err != nil {
// 			c.JSON(500, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(200, gin.H{"room": room})
// 	})
// }
