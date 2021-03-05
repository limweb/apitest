package main

import (
	_ "apitest/docs"
	"apitest/handler"
	"log"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title ApiTest
// @version 2.0
// @description.markdown
// @termsOfService http://somewhere.com/

// @contact.name API Support
// @contact.url http://somewhere.com/support
// @contact.email support@somewhere.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	r.GET("/healthcheck", handler.HealthCheckHandler)

	v1 := r.Group("/api/v1")
	{
		customers := v1.Group("/customers")
		{
			customersHandler := handler.CustomerHandler{}
			customers.GET(":id", customersHandler.GetCustomer)
			customers.GET("", customersHandler.ListCustomers)
			customers.POST("", customersHandler.CreateCustomer)
			customers.DELETE(":id", customersHandler.DeleteCustomer)
			customers.PATCH(":id", customersHandler.UpdateCustomer)
		}
	}
	r.GET("/apidoc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// log.Fatal((r.Run(":8080")))
	log.Fatal((r.Run()))
}
