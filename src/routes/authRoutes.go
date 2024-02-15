package routes

import (
	"EmployeeManagementTool/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/employee/signup", controllers.SignupManager())
	router.POST("/employee/login", controllers.LoginManager())
}
