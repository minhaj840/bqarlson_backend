package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/minhaj840/bqarlson_backend/controller"
	"github.com/minhaj840/bqarlson_backend/middleware"

)

func SetupRoutes(c *gin.Engine) {

	c.POST("/login", controller.Login)
	c.POST("/register", controller.Register)

	private := c.Group("/private")

	private.Use(middleware.Authenticate)

	private.GET("/refreshtoken", controller.RefreshToken)
	private.GET("/profile", controller.Profile)

}
