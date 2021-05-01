package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

func EmailConfirmation(c *gin.Context) {
	query := c.Request.URL.Query()
	email := query["email"]
	token := query["token"]
	c.JSON(http.StatusOK, gin.H{"email": email, "token": token, "message": "Registration Success"})
	user, _ := database.GetRegistration(email[0], token[0])
	c.JSON(http.StatusOK, user)
}
