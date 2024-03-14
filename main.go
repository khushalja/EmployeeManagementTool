package main

import (

	// "EmployeeManagementTool/src/model"
	"EmployeeManagementTool/src/configs"
	// "EmployeeManagementTool/src/mongodb"
	"EmployeeManagementTool/src/routes"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("Emp management tool")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	// restapi.GinConnection()
	// uri := "mongodb://localhost:27017"
	// client, ctx, cancel, err := mongodb.ConnectDb(uri)
	client, ctx, cancel, err := configs.ConnectDb(configs.EnvMongoURI())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Successfull!!")
	// defer mongodb.Close(client, ctx, cancel)
	defer configs.Close(client, ctx, cancel)

	router := gin.Default()
	routes.AuthRoutes(router)
	routes.GinConnection(router)

}
