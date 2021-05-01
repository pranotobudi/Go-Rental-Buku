package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

func ListAllBooks(c *gin.Context) {
	// r.GET("/books", router.ListAllBooks)

	books, err := database.ListAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, books)

}
func ListBook(c *gin.Context) {
	// r.GET("/books/:id", router.ListBook)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	fmt.Println("ListCategory ERROR: ", err)

	book, err := database.ListBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)

}
func AddBook(c *gin.Context) {
	// r.POST("/books", router.AddBook)
	id_category, _ := strconv.Atoi(c.PostForm("id_category"))
	title := c.PostForm("title")
	description := c.PostForm("description")
	author := c.PostForm("author")
	year, _ := strconv.Atoi(c.PostForm("year"))
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	err := database.AddBook(id_category, title, description, author, year, stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Add Book Success")
}
func UpdateBook(c *gin.Context) {
	// r.PUT("/books/:id", router.UpdateBook)
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	id_category, _ := strconv.Atoi(c.PostForm("id_category"))
	title := c.PostForm("title")
	description := c.PostForm("description")
	author := c.PostForm("author")
	year, _ := strconv.Atoi(c.PostForm("year"))
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	err := database.UpdateBook(id, id_category, title, description, author, year, stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Update Book Success id: "+idStr)
}
func DeleteAllBooks(c *gin.Context) {
	// r.DELETE("/books", router.DeleteAllBooks)
	err := database.DeleteAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete All Books Success")
}
func DeleteBook(c *gin.Context) {
	// r.DELETE("/books/:id", router.DeleteBook)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := database.DeleteBook(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete Book Success, id: "+idStr)
}
