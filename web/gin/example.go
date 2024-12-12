package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
var db = make(map[string]string)

func populateDb(){
	db["kitty"] = "hello"
	db["homer"] = "hello"
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// ping test
	r.GET("/ping", func(c *gin.Context){
		c.String(http.StatusOK, "pong")
	})
}

func main() {
	
	fmt.Printf("%+v", db)
}