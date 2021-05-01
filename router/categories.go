package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

//CRUD Category
func ListAllCategories(c *gin.Context) {
	//	r.GET("/categories", router.ListAllCategories)

	categories, err := database.ListAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, categories)

}
func ListCategory(c *gin.Context) {
	// r.GET("/categories/:id", router.ListCategory)
	// query := c.Request.URL.Query()
	// idStr := query["id"]
	// id, err := strconv.Atoi(idStr[0])
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	category, err := database.ListCategory(id)
	fmt.Println("ListCategory ERROR: ", err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, category)

}
func AddCategory(c *gin.Context) {
	// r.POST("/categories", router.AddCategory)
	name := c.PostForm("name")
	err := database.AddCategory(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Add Category Success")
}
func UpdateCategory(c *gin.Context) {
	// r.PUT("/categories/:id", router.UpdateCategory)
	// query := c.Request.URL.Query()
	// idStr := query["id"]
	// id, _ := strconv.Atoi(idStr[0])
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	name := c.PostForm("name")
	err := database.UpdateCategory(id, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Update Category Success id: "+idStr)
}

func DeleteAllCategories(c *gin.Context) {
	// r.DELETE("/categories", router.DeleteAllCategories)
	err := database.DeleteAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete All Categories Success")
}
func DeleteCategory(c *gin.Context) {
	// r.DELETE("/categories/:id", router.DeleteCategory)

	// query := c.Request.URL.Query()
	// idStr := query["id"]
	// id, _ := strconv.Atoi(idStr[0])
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := database.DeleteCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete Category Success, id: "+idStr)
}

func GetCategoryBasedBook(c *gin.Context) {
	id_category := c.PostForm("id_category")
	categoryInt, _ := strconv.Atoi(id_category)
	books, err := database.GetCategoryBasedBookList(categoryInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, books)
	}
	c.JSON(http.StatusOK, books)
}
