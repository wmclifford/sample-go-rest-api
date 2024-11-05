package tests

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(models ...interface{}) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	log.Printf("db: %v", db)
	if err != nil {
		log.Fatalf("failed to connect to the test database: %v", err)
	}
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("failed to migrate the test database: %v", err)
			return nil
		}
	}
	return db
}
