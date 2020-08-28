package main

import (
	"github.com/ShAlireza/gin_tutorial/controllers"
	"github.com/ShAlireza/gin_tutorial/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello world!"})
}

func testHandler(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/books", controllers.FindBooks)


	r.Run()

}
