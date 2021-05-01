package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

func UserLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if IsMember(email, password) {
		user, _ := database.GetUserByEmail(email)
		token, _ := GenerateJWT(email, user.Role)
		c.JSON(http.StatusOK, gin.H{"jwt_token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Credential Not Found, Please Login First"})
	}
}
