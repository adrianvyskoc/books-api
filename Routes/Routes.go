package Routes

import (
	"books/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	booksRoutes := r.Group("/api")
	{
		booksRoutes.GET("books", Controllers.GetBooks)
		booksRoutes.POST("books", Controllers.CreateBook)
		//booksRoutes.GET("books/:id", Controllers.GetBookByID)
  	//booksRoutes.PUT("books/:id", Controllers.UpdateBook)
		booksRoutes.DELETE("books/:id", Controllers.DeleteBook)
	}
	return r
}