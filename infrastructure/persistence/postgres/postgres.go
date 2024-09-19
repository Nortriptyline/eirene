package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabaseConnection() error {
	// Retrieve the database connection parameters from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construct the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the database connection
	var err error
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&models.BankGormModel{},
		&models.BankAccountGormModel{},
		&models.TransactionGormModel{},
	)

	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() (*gorm.DB, error) {
	var err error

	if db == nil {
		err = InitDatabaseConnection()
	}

	return db, err
}

// CloseDatabaseConnection closes the database connection
func CloseDatabaseConnection() {
	if db == nil {
		log.Fatalf("Error closing the database: database connection is nil")
	}

	conn, err := db.DB()
	if err != nil {
		log.Fatalf("Error closing the database: %v", err)
	}

	err = conn.Close()
	if err != nil {
		log.Fatalf("Error closing the database: %v", err)
	}
}
