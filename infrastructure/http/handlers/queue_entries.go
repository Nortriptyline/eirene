package handlers

// import (
// 	"github.com/Nortriptyline/Eirene/application/commands"
// 	"github.com/Nortriptyline/Eirene/application/dto"
// 	"github.com/Nortriptyline/Eirene/application/handlers/commandhandlers"
// 	"github.com/Nortriptyline/Eirene/domain/service"
// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// func CreateQueueEntry(router *gin.RouterGroup, commandHandler *commandhandlers.CreateQueueEntryCommandHandler) {
// 	router.POST("/:queue_id/entries", func(c *gin.Context) {
// 		// Make sur the QueueID is a valid UUID

// 		qid, err := uuid.Parse(c.Param("queue_id"))

// 		if err != nil {
// 			c.JSON(400, gin.H{"error": "Invalid queue ID"})
// 			return
// 		}

// 		cmd := commands.CreateQueueEntryCommand{
// 			QueueID:     qid,
// 			ClientToken: c.MustGet("clientToken").(string),
// 		}

// 		entry, err := commandHandler.Handle(cmd)

// 		// Handle the error
// 		if err != nil {
// 			switch err {
// 			case service.ErrQueueNotFound:
// 				c.JSON(404, gin.H{"error": "Queue not found"})
// 			case service.ErrInternal:
// 				c.JSON(500, gin.H{"error": "Internal server error"})
// 			default:
// 				c.JSON(500, gin.H{"error": err.Error()})
// 			}
// 			return
// 		}

// 		c.JSON(201, gin.H{"entry": entry})
// 	})
// }

// func CreateLeaveQueue(router *gin.RouterGroup, commandHandler *commandhandlers.LeaveQueueCommandHandler) {
// 	router.PATCH("/:queue_id/leave", func(c *gin.Context) {

// 		qid, err := uuid.Parse(c.Param("queue_id"))

// 		if err != nil {
// 			c.JSON(400, gin.H{"error": "Invalid queue ID"})
// 			return
// 		}

// 		cmd := commands.LeaveQueueCommand{
// 			QueueID:     qid,
// 			ClientToken: c.MustGet("clientToken").(string),
// 		}

// 		entry, err := commandHandler.Handle(cmd)

// 		if err != nil {
// 			switch err {
// 			case service.ErrQueueNotFound:
// 				c.JSON(404, gin.H{"error": err.Error()})
// 			case service.ErrNotInQueue:
// 				c.JSON(404, gin.H{"error": err.Error()})
// 			default:
// 				c.JSON(500, gin.H{"error": err.Error()})
// 			}
// 			return
// 		}

// 		entryDto := dto.ToQueueEntryDTO(entry)
// 		c.JSON(201, gin.H{"entry": entryDto})
// 	})
// }

// func CreateHoldQueueEntry(router *gin.RouterGroup, commandHandler *commandhandlers.HoldQueueEntryCommandHandler) {
// 	router.PATCH("/:queue_id/hold", func(c *gin.Context) {
// 		qid, err := uuid.Parse(c.Param("queue_id"))

// 		if err != nil {
// 			c.JSON(400, gin.H{"error": "Invalid queue ID"})
// 			return
// 		}

// 		cmd := commands.HoldQueueEntryCommand{
// 			QueueID:     qid,
// 			ClientToken: c.MustGet("clientToken").(string),
// 		}

// 		entry, err := commandHandler.Handle(cmd)

// 		if err != nil {

// 			switch err {
// 			case service.ErrQueueNotFound:
// 				c.JSON(404, gin.H{"error": err.Error()})
// 			case service.ErrNotInQueue:
// 				c.JSON(404, gin.H{"error": err.Error()})
// 			default:
// 				c.JSON(500, gin.H{"error": err.Error()})
// 			}

// 			return
// 		}

// 		entryDto := dto.ToQueueEntryDTO(entry)

// 		c.JSON(201, gin.H{"entry": entryDto})
// 	})
// }
