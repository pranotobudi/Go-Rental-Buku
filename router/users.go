package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

func ListAllUsers(c *gin.Context) {
	// r.GET("/users", router.ListAllUsers)

	books, err := database.ListAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, books)

}
func ListUser(c *gin.Context) {
	// r.GET("/users/:id", router.ListUser)
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	fmt.Println("ListCategory ERROR: ", err)

	book, err := database.ListUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)

}
func AddUser(c *gin.Context) {
	// r.POST("/users", router.AddUser)

	name := c.PostForm("name")
	address := c.PostForm("address")
	photo := c.PostForm("photo")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role, _ := strconv.Atoi(c.PostForm("role"))
	err := database.AddUser(name, address, photo, email, password, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Add Book Success")
}
func UpdateUser(c *gin.Context) {
	// r.PUT("/users/:id", router.UpdateUser)
	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	name := c.PostForm("name")
	address := c.PostForm("address")
	photo := c.PostForm("photo")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role, _ := strconv.Atoi(c.PostForm("role"))

	err := database.UpdateUser(id, name, address, photo, email, password, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Update Book Success id: "+idStr)
}

func DeleteAllUsers(c *gin.Context) {
	// r.DELETE("/users", router.DeleteAllUsers)

	err := database.DeleteAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete All Books Success")
}
func DeleteUser(c *gin.Context) {
	// r.DELETE("/users/:id", router.DeleteUser)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := database.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete Book Success, id: "+idStr)
}
