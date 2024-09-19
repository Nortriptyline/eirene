package main

import (
	"github.com/Nortriptyline/Eirene/application"
	"github.com/Nortriptyline/Eirene/infrastructure/http"
	"github.com/Nortriptyline/Eirene/infrastructure/logging"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/postgres"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	// Once the .env file is loaded, load the logstash logger
	// It is the one that will be used for the rest of the application
	sugar := logging.GetLogger()

	if sugar == nil {
		panic("Error loading logger")
	}

	db, err := postgres.GetDatabase()
	if err != nil {
		panic("Error loading database")
	}

	defer logging.CloseLogstashConnection()
	defer postgres.CloseDatabaseConnection()
	defer sugar.Sync()

	// eventBus := eventbus.NewEventBus()

	repositories := application.InitializeRepositories(db)
	services := application.InitializeServices(repositories, sugar)

	// Subscribing to the event
	// eventBus.Subscribe("events.QueueEntryUpdatedEvent", queueService.HandleQueueEntryUpdated)
	// eventBus.Subscribe("domain.QueueLeftEvent", queueService.HandleQueueEntryLeftEvent)

	handlers := &application.Handlers{
		Commands: application.InitializeCommands(services),
		Queries:  application.InitializeQueries(),
	}

	router := gin.Default()

	http.RegisterRoutes(router, handlers)
	http.RegisterGraphQlRoutes(router)

	router.Run(":8080")
}
