package postgres

import (
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.Exec("PRAGMA foreign_keys = ON")
	db.AutoMigrate(
		&models.BankGormModel{},
		&models.BankAccountGormModel{},
		&models.TransactionGormModel{},
	)
	return db
}
