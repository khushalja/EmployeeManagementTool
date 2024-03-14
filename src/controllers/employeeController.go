package controllers

import (
	"EmployeeManagementTool/src/configs"
	"EmployeeManagementTool/src/helper"
	"EmployeeManagementTool/src/model"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoginManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)

		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvMongoURI()))
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		collection := client.Database(configs.EnvDatabase()).Collection(configs.EnvCollection())

		var emp model.Employee
		var existingEmp model.Employee

		if err := c.ShouldBindJSON(&emp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "failed to bind json"})
			return
		}

		err = collection.FindOne(ctx, bson.D{{Key: "employeeid", Value: emp.EmployeeId}}).Decode(&existingEmp)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "failed to retieve employee"})
			return
		}

		if emp.EmployeeId != existingEmp.EmployeeId || emp.Password != existingEmp.Password {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide correct EmployeeId or Password"})
			return
		}
		token, refreshedtoken, err := helper.GenerateAllTokens(existingEmp.EmployeeName)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		existingEmp.Token = token
		existingEmp.RefreshedToken = refreshedtoken

		c.JSON(http.StatusOK, gin.H{"token": existingEmp.Token, "Employee Logged in successfully": existingEmp.EmployeeName})

	}
}

func GetAllEmployees() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)

		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvMongoURI()))
		if err != nil {
			log.Fatal(err)
		}
		collection := client.Database(configs.EnvDatabase()).Collection(configs.EnvCollection())
		var emps []model.Employee
		// var returnemp []model.Employee
		filter := bson.D{{}}
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			log.Fatal(err)
		}
		if err := cursor.All(ctx, &emps); err != nil {
			log.Panic(err)
		}
		// for _, result := range emps{
		// 	res, _:= bson.MarshalExtJSON(result, false , false)
		// 	returnemp= append(returnemp, res)
		// }

		fmt.Println(emps) //to delete
		c.JSON(http.StatusOK, emps)
	}
}
func GetEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		empid := c.Param("empid")
		ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)

		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvMongoURI()))
		if err != nil {
			log.Fatal(err)
		}
		collection := client.Database(configs.EnvDatabase()).Collection(configs.EnvCollection())
		var emp model.Employee
		filter := bson.D{{Key: "emplyoid", Value: empid}}
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			log.Fatal(err)
		}
		if err := cursor.All(ctx, &emp); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, emp)
	}
}

func CreateEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		var emp model.Employee
		if err := c.ShouldBindJSON(&emp); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)

		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvMongoURI()))
		if err != nil {
			log.Fatal(err)
		}
		collection := client.Database(configs.EnvDatabase()).Collection(configs.EnvCollection())

		result, err := collection.InsertOne(ctx, emp)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, result)
	}
}

func UpdateEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		empid := c.Param("empid")
		var emp model.Employee
		if err := c.BindJSON(&emp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)

		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvMongoURI()))
		if err != nil {
			log.Fatal(err)
		}
		collection := client.Database(configs.EnvDatabase()).Collection(configs.EnvCollection())

		filter := bson.D{{Key: "employeeid", Value: empid}}
		update := bson.D{{Key: "$set", Value: emp}}
		result, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.ModifiedCount)

		var emp1 model.Employee

		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			log.Fatal(err)
		}
		if err := cursor.All(ctx, &emp); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, emp1)
	}
}

func DeleteEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		empid := c.Param("empid")
		ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvMongoURI()))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection := client.Database(configs.EnvDatabase()).Collection(configs.EnvCollection())

		result, err := collection.DeleteOne(ctx, bson.D{{Key: "employeeid", Value: empid}})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)

	}
}
