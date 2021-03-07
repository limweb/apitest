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
	r.GET("/healthcheck", controllers.HealthCheckHandler)
	controllers.SetupLoginRoutes(r)
	controllers.SetupAuthRoutes(r)
	v1 := r.Group("/api/v1")
	{
		customers := v1.Group("/customers")
		{
			customersHandler := controllers.CustomerHandler{}
			customers.GET(":id", customersHandler.GetCustomer)
			customers.GET("", customersHandler.ListCustomers)
			customers.POST("", customersHandler.CreateCustomer)
			customers.DELETE(":id", customersHandler.DeleteCustomer)
			customers.PATCH(":id", customersHandler.UpdateCustomer)
		}
	}
	controllers.SetupBookRoutes(v1)
	controllers.SetupTestRoutes(v1)
	r.GET("/apidoc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// log.Fatal((r.Run(":8080")))
	_ = port
	// log.Fatal((r.Run(port)))
	log.Fatal((r.Run()))
}
