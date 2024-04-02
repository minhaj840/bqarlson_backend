package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/minhaj840/bqarlson_backend/database"
	"github.com/minhaj840/bqarlson_backend/routes"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error in loading env file.")
	}

	database.ConnectDB()
}

func main() {

	// Close the db connection using defer clause
	sqlDb, err := database.DBConn.DB()

	if err != nil {
		log.Println("Error in getting db conn.")
	}

	defer sqlDb.Close()

	port := os.Getenv("port")

	if port == "" {
		port = "8001"
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// // Add middleware
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
	// 	AllowHeaders: []string{"Origin", "Auth-token", "token", "Content-type"},
	// }))
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, PATCH, GET, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Auth-token, token, Content-type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	routes.SetupRoutes(router)

	log.Fatal(router.Run(":" + port))
}
