package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

func TransactionList(c *gin.Context) {
	email := c.PostForm("email")
	user, _ := database.GetUserByEmail(email)
	transactions, err := database.GetTransactionList(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, transactions)
	}
	c.JSON(http.StatusOK, transactions)
}
