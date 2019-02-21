package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	app.GET("/api/hello", handler_hello)
}

func handler_hello(c *gin.Context) {

}
