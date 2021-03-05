package main

import (
	"apitest/db"
	"apitest/model"
	"apitest/webserver"
	"flag"
	"log"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "8080", "web port 8080")
}

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
	flag.Parse()
	log.Printf("----------------------  Web Server Api  V 0.0.1   --------------------------")
	log.Printf("----------------------วิธีใช้: server.exe -port 8080 --------------------------")
	port = ":" + port
	log.Printf("-------------------Starting with WebServer on Port %s --------------------", port)
	db.SetupDB()
	p1 := model.Person{FirstName: "John", LastName: "Doe"}
	p2 := model.Person{FirstName: "Jane", LastName: "Smith"}
	db.GetDB().Create(&p1)
	db.GetDB().Create(&p2)
	webserver.StartWeb(port)
}
