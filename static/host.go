package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.Static("/static", "static")

	router.Run(":8090")
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// e.Use(middleware.CORS())

	// e.Static("/static", "static")

	// e.Logger.Fatal(e.Start(":8090"))
}
