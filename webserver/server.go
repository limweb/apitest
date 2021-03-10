package webserver

import (
	"apitest/controllers"
	_ "apitest/docs"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func StartWeb(port string) {
	r := gin.Default()

	// Set up CORS middleware options
	config := cors.Config{
		Origins:        "*",
		RequestHeaders: "Origin, Authorization, Content-Type",

		Methods:         "GET, POST, PUT, DELETE",
		Credentials:     true,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}

	// Apply the middleware to the router (works on groups too)
	r.Use(cors.Middleware(config))
	r.Static("/images", "./publics/upload/images")
	r.GET("/", controllers.Message)
	r.GET("/ping", controllers.Ping)
	r.GET("/test", controllers.Test)
	r.GET("/healthcheck", controllers.HealthCheckHandler)
	r.GET("/apidoc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	controllers.SetupLoginRoutes(r)
	controllers.SetupAuthRoutes(r)
	v1 := r.Group("/api/v1")
	controllers.SetupBookRoutes(v1)
	// log.Fatal((r.Run(":8080")))
	_ = port
	// log.Fatal((r.Run(port)))
	log.Fatal((r.Run()))
}
