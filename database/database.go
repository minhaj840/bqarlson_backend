package database

import (
	"log"

	"github.com/minhaj840/bqarlson_backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

// func ConnectDB() {
// 	dsn := os.Getenv("dsn")

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Error),
// 	})

// 	if err != nil {
// 		panic("Error in database connection.")
// 	}
// 	log.Println("Databse conncection successful.")

// 	db.AutoMigrate(new(model.User))

// 	DBConn = db
// }

func ConnectDB() {
	// dsn := os.Getenv("POSTGRES_DSN") // Update the environment variable name
	dsn := "host=localhost port=5432 user=api_user password=your_password dbname=your_database sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Error in database connection.")
	}
	log.Println("Database connection successful.")

	db.AutoMigrate(&model.User{}) // Update the model migration

	DBConn = db
}
