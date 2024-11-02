package db

import (
	"log"

	"backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(dbConfig config.DBConfig) *gorm.DB {
	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	return db
}
