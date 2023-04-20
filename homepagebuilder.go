package main

import (
	//"html/template"

	"time"

	_ "os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/vadimzharov/homepagebuilder/pkg/handlers"
	"github.com/vadimzharov/homepagebuilder/pkg/pagebuilder"
	_ "github.com/vadimzharov/homepagebuilder/pkg/utils"
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

	router.GET("/api/regenerateall", handlers.RegenerateAll)
	router.GET("/api/regeneratemain", handlers.RegenerateMain)
	router.GET("/api/regeneratepanels", handlers.RegeneratePanels)
	router.GET("/api/regenerateapps", handlers.RegenerateAppsDescription)
	router.Run(":8080")

}
