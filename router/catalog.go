package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

func GetCatalog(c *gin.Context) {

	books, err := database.GetAllBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, books)
	}
	c.JSON(http.StatusOK, books)
}
