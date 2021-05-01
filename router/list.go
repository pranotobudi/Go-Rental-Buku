package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

func UserList(c *gin.Context) {

	users, err := database.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, users)
	}
	c.JSON(http.StatusOK, users)
}
