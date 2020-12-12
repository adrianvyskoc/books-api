package Controllers

import (
	"books/Models"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var book []Models.Book
	err := Models.GetAllBooks(&book)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func CreateBook(c *gin.Context) {
	var book Models.Book
	c.BindJSON(&book)
	err := Models.CreateBook(&book)
	if err != nil {
	 fmt.Println(err.Error())
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, book)
	}
}

func DeleteBook(c *gin.Context) {
	id := c.Params.ByName("id")
	err := Models.DeleteBookByID(id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, gin.H{"result": "document with id " + id + " is deleted"})
	}
 }