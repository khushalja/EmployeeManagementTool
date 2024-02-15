package middleware

import (
	"EmployeeManagementTool/src/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no authorization header provided"})
			c.Abort()
			return
		}
		// splitToken := strings.Split(clientToken, " ")
		// if len(splitToken) == 2 {
		// 	clientToken = splitToken[1]
		// }
		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("employee_name", claims.EmployeeName)
		// c.Set("employee_id", claims.EmployeeId)
		// c.Set("employee_type", claims.EmployeeType)
		c.Next()
	}
}
