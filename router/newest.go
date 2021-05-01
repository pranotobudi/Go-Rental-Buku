package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

func GetNewest(c *gin.Context) {
	var maxShowBook int = 20
	var finalBooks []database.Book
	books, err := database.GetNewestList()
	var counter int
	for _, book := range books {
		counter++
		finalBooks = append(finalBooks, book)
		if counter > maxShowBook {
			break
		}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, finalBooks)
	}
	c.JSON(http.StatusOK, finalBooks)
}
