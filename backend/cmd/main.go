package main

import (
	"Chequeandome-Vision-Demo/docs"
	"Chequeandome-Vision-Demo/internal/configs"
	_ "Chequeandome-Vision-Demo/internal/configs"
	middlewares "Chequeandome-Vision-Demo/internal/middleware"
	"Chequeandome-Vision-Demo/internal/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strings"
)

func CORS() gin.HandlerFunc {
	// Reference: https://github.com/gin-contrib/cors/issues/29
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORS())
	r.Use(middlewares.BodySizeLimiter())

	var host string
	if strings.Contains(configs.URL, "https://") {
		host = strings.Split(configs.URL, "https://")[1]
	} else {
		host = strings.Split(configs.URL, "http://")[1]
	}

	// Swagger
	docs.SwaggerInfo.Title = "File server"
	docs.SwaggerInfo.Description = "Simple file server for storing files"
	docs.SwaggerInfo.Version = "0.1.0"
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = "/"

	// Routes
	//urlSwagger := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.Routes(r)

	err := r.Run() // listen and serve on
	if err != nil {
		panic(err)
	}
}
