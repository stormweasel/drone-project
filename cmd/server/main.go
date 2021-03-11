package main

import (
	"log"
	"pig/cmd/login"
	"pig/cmd/register"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../../web", true)))

	router.POST("/register", register.RegisterTheUser)
	router.POST("/login", login.LoginFunction)
	log.Fatal(router.Run(":8080"))
}
