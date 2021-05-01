package router

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

//CRUD Category
func ListAllLoans(c *gin.Context) {
	// r.GET("/loans", router.ListAllLoans)
	loans, err := database.ListAllLoans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, loans)

}
func ListLoan(c *gin.Context) {
	// r.GET("/loans/:id", router.ListLoan)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	loan, err := database.ListLoan(id)
	fmt.Println("ListLoan ERROR: ", err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, loan)

}

func AddLoan(c *gin.Context) {
	// r.POST("/loans", router.AddLoan)
	idUser, _ := strconv.Atoi(c.PostForm("id_user"))
	idBook, _ := strconv.Atoi(c.PostForm("id_book"))
	borrowedDate := time.Now()
	dueDate := time.Now().AddDate(0, 0, 3)
	returnDate := time.Now()
	err := database.AddLoan(uint(idUser), uint(idBook), borrowedDate, dueDate, returnDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Add Loan Success")
}
func UpdateLoan(c *gin.Context) {
	// r.PUT("/loans/:id", router.UpdateLoan)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	idUser, _ := strconv.Atoi(c.PostForm("id_user"))
	idBook, _ := strconv.Atoi(c.PostForm("id_book"))
	borrowedDate := time.Now()
	dueDate := time.Now().AddDate(0, 0, 3)
	returnDate := time.Now()

	err := database.UpdateLoan(id, uint(idUser), uint(idBook), borrowedDate, dueDate, returnDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Update Loan Success id: "+idStr)
}

func DeleteAllLoans(c *gin.Context) {
	// r.DELETE("/loans", router.DeleteAllLoans)
	err := database.DeleteAllLoans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete All Categories Success")
}
func DeleteLoan(c *gin.Context) {
	// r.DELETE("/loans/:id", router.DeleteLoan)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := database.DeleteLoan(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete Category Success, id: "+idStr)
}
