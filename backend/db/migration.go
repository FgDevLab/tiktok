package db

import (
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	fmt.Println("Starting database migration...")

	if err := db.AutoMigrate(); err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	fmt.Println("Database migration completed successfully.")
	return nil
}
