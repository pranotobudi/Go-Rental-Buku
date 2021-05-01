package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

func ListAllFinePayments(c *gin.Context) {
	// r.GET("/fine-payments", router.ListAllFinePayments)
	books, err := database.ListAllFinePayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, books)

}
func ListFinePayment(c *gin.Context) {

	// r.GET("/fine-payments/:id", router.ListFinePayment)
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	fmt.Println("ListCategory ERROR: ", err)

	book, err := database.ListFinePayment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)

}
func AddFinePayment(c *gin.Context) {
	// r.POST("/fine-payments", router.AddFinePayment)
	receipt := c.PostForm("receipt")
	amount, _ := strconv.Atoi(c.PostForm("amount"))
	idLoan, _ := strconv.Atoi(c.PostForm("idLoan"))

	err := database.AddFinePayment(receipt, int64(amount), uint(idLoan))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Add FinePayment Success")
}
func UpdateFinePayment(c *gin.Context) {

	// r.PUT("/fine-payments/:id", router.UpdateFinePayment)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(c.PostForm(idStr))
	receipt := c.PostForm("receipt")
	amount, _ := strconv.Atoi(c.PostForm("amount"))
	idLoan, _ := strconv.Atoi(c.PostForm("idLoan"))

	err := database.UpdateFinePayment(id, receipt, int64(amount), uint(idLoan))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Update FinePayment Success id: "+idStr)
}

func DeleteAllFinePayments(c *gin.Context) {
	// r.DELETE("/fine-payments", router.DeleteAllFinePayments)

	err := database.DeleteAllFinePayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete All FinePayments Success")
}
func DeleteFinePayment(c *gin.Context) {
	// r.DELETE("/fine-payments/:id", router.DeleteFinePayment)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := database.DeleteFinePayment(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete FinePayment Success, id: "+idStr)
}
