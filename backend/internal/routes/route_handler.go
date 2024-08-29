package routes

import (
	"Chequeandome-Vision-Demo/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":       "File Service is up and running!",
			"documentation": "https://redes-markalbrand56.koyeb.app/files/documentation/index.html",
		})
	})

	r.POST("/", controllers.Upload)
	r.GET("/:file", controllers.GetFile)
}
