package main

import (
	//"html/template"

	"time"

	//"net/http"
	_ "os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/vadimzharov/homepagebuilder/pkg/pagebuilder"
)

func main() {

	pagebuilder.BuildNewPage()

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./assets", false)))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/api/regenerateall", regenerateAll)
	router.Run(":8080")

}

func regenerateAll(c *gin.Context) {
	pagebuilder.BuildNewPage()
	return
}
