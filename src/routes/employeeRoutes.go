package routes

import (
	"EmployeeManagementTool/src/controllers"
	"EmployeeManagementTool/src/middleware"

	"github.com/gin-gonic/gin"
)

func GinConnection(r *gin.Engine) {
	// router := gin.Default()
	router := r.Group("/employee")
	router.Use(middleware.Authenticate())
	{
		router.GET("/employeedetails", controllers.GetAllEmployees())
		router.GET("/employeedetails/:empid", controllers.GetEmployee())
		router.POST("/employeedetails/create", controllers.CreateEmployee())
		router.PUT("/employeedetails/:empid", controllers.UpdateEmployee())
		router.DELETE("/employeedetails/:empid", controllers.DeleteEmployee())
	}
	r.Run(":6000")
}
