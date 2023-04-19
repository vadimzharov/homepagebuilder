package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vadimzharov/homepagebuilder/pkg/pagebuilder"
	"github.com/vadimzharov/homepagebuilder/pkg/utils"
)

func RegenerateAll(c *gin.Context) {

	utils.EraseAllSourceFiles()
	go pagebuilder.BuildNewPage()
	return
}

func RegenerateMain(c *gin.Context) {
	utils.EraseSourceFile("index.html")
	go pagebuilder.BuildNewPage()
	return
}

func RegeneratePanels(c *gin.Context) {
	utils.EraseSourceFile("panel-code.tmpl")
	utils.EraseSourceFile("panels-code.html")
	go pagebuilder.BuildNewPage()
	return
}

func RegenerateAppsDescription(c *gin.Context) {
	utils.EraseSourceFile("appsdescription.json")
	utils.EraseSourceFile("panels-code.html")
	go pagebuilder.BuildNewPage()
	return
}
