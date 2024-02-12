package routes

import (
	"EmployeeManagementTool/src/controllers"

	"github.com/gin-gonic/gin"
)

func GinConnection(router *gin.Engine) {
	// router := gin.Default()

	router.GET("/employeedetails", controllers.GetAllEmployees())
	router.GET("/employeedetails/:empid", controllers.GetEmployee())
	router.POST("/employeedetails/create", controllers.CreateEmployee())
	router.PUT("/employeedetails/:empid", controllers.UpdateEmployee())
	router.DELETE("/employeedetails/:empid", controllers.DeleteEmployee())

	router.Run(":6000")
}
