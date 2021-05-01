package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

func BookRentList(c *gin.Context) {
	email := c.PostForm("email")
	user, _ := database.GetUserByEmail(email)
	books, err := database.GetBookRent(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, books)
	}
	c.JSON(http.StatusOK, books)
}
