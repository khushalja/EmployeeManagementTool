package main

import (

	// "EmployeeManagementTool/src/model"
	"EmployeeManagementTool/src/mongodb"
	"EmployeeManagementTool/src/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("Emp management tool")
	// restapi.GinConnection()
	uri := "mongodb://localhost:27017"
	client, ctx, cancel, err := mongodb.ConnectDb(uri)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Successfull!!")
	defer mongodb.Close(client, ctx, cancel)
	// emp1 := model.Employee{
	// 	EmployeeId:   1113,
	// 	EmployeeName: "MC guru",
	// 	Contact: model.Contact{
	// 		PhoneNo: 2000000000,
	// 		Address: "chal chal",
	// 		EmailId: "abhiKaamNahiHai@faltu.com",
	// 	},
	// 	JobTitle:   "Systems Engineer",
	// 	Department: "Music indutry",
	// 	Salary:     10.00,
	// }
	// emp2 := model.Employee{
	// 	EmployeeId:   1112,
	// 	EmployeeName: "MC Sher",
	// 	Contact: model.Contact{
	// 		PhoneNo: 1000000001,
	// 		Address: "jara nach ke dikha",
	// 		EmailId: "Mcsher@faltu.com",
	// 	},
	// 	JobTitle:   "Systems Engineer",
	// 	Department: "Music indutry",
	// 	Salary:     12.00,
	// }
	// var empData []interface{}

	// empData = append(empData, emp1, emp2)
	// result, err := mongodb.InsertOne(client, ctx, "EmployeeManagement", "EmployeeDetails", emp1)
	// result, err := mongodb.InsertMany(client, ctx, "EmployeeManagement", "EmployeeDetails", empData)

	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("Displaying Id of Inserted document")
	// fmt.Println(result.InsertedIDs...)
	// filter := bson.D{{Key: "employeeid", Value: bson.D{{Key: "$eq", Value: 1113}}}}
	// fmt.Println(filter)
	// updateData := bson.D{{Key: "$set", Value: bson.D{{Key: "salary", Value: 120.00}}}}
	// result, err := mongodb.UpdateOne(client, ctx, "EmployeeManagement", "EmployeeDetails", filter, updateData)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(result.ModifiedCount)
	router := gin.Default()
	routes.AuthRoutes(router)
	routes.GinConnection(router)

}
