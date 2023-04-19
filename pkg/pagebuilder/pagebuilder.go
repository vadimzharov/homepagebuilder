package pagebuilder

import (
	"bytes"
	"fmt"
	_ "html/template"
	"log"
	"strings"
	"text/template"

	"encoding/json"
	_ "encoding/json"

	"github.com/vadimzharov/homepagebuilder/pkg/chatgpt"
	"github.com/vadimzharov/homepagebuilder/pkg/config"
	"github.com/vadimzharov/homepagebuilder/pkg/utils"
)

type Applications struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var DBConfig *config.DashboardConfig

func GeneratePanelsConfig(apps []Applications) {
	DBConfig = config.ReadConfig()

	for i := range DBConfig.Panels {

		log.Println("Creating config for panel named", DBConfig.Panels[i].Name)
		for j := range apps {
			if DBConfig.Panels[i].Name == apps[j].Name {
				DBConfig.Panels[i].Description = apps[j].Description
			}
		}

	}

	log.Println("======== Panels configuration struct ==============")
	log.Println(DBConfig.Panels)

}

func CombinePanelsCode(panelcode string) string {

	var panelscode strings.Builder

	var paneltmplexec bytes.Buffer

	for i := range DBConfig.Panels {

		paneltmpl, err := template.New("paneltemplate").Parse(panelcode)

		if err != nil {
			panic(err)
		}

		err = paneltmpl.Execute(&paneltmplexec, DBConfig.Panels[i])

		if err != nil {
			panic(err)
		}

		panelscode.WriteString(paneltmplexec.String())
		panelscode.WriteString("\n")

		paneltmplexec.Reset()

	}

	return panelscode.String()
}

func BuildNewPage() {

	dbconfig := config.ReadConfig()

	log.Println("===== Readed config =====")
	log.Println(dbconfig.Panels)

	if utils.SourceFileExists("index.html") == false {

		landingpagecode := chatgpt.QueryGPT(dbconfig.ChatGPTAPIKey, mainpagequery)

		utils.WriteToSourceFile(landingpagecode, "index.html")

	}

	var applications []Applications

	var appsdescription []byte

	if utils.SourceFileExists("appsdescription.json") == false {

		var appsquery strings.Builder

		var appelement bytes.Buffer

		appsquery.WriteString(appsreqs)

		for i := range dbconfig.Panels {

			appstempl, err := template.New("appstemplate").Parse(apptemplate)
			if err != nil {
				panic(err)
			}

			err = appstempl.Execute(&appelement, dbconfig.Panels[i])

			if err != nil {
				panic(err)
			}

			appsquery.WriteString(appelement.String())

			appelement.Reset()

			// log.Println(appelement.String())

		}

		log.Println(appsquery.String())

		appsdescription = []byte(chatgpt.QueryGPT(dbconfig.ChatGPTAPIKey, appsquery.String()))

		log.Println(appsdescription)

		utils.WriteToSourceFile(string(appsdescription), "appsdescription.json")

	} else {
		appsdescription = utils.ReadFromSourceFile("appsdescription.json")
	}

	err := json.Unmarshal([]byte(appsdescription), &applications)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range applications {
		log.Println(applications[i].Name)

	}

	var panelcode string

	if utils.SourceFileExists("panel-code.tmpl") == false {

		panelcode = chatgpt.QueryGPT(dbconfig.ChatGPTAPIKey, panelquery)

		utils.WriteToSourceFile(panelcode, "panel-code.tmpl")

	} else {
		panelcode = string(utils.ReadFromSourceFile("panel-code.tmpl"))
	}

	GeneratePanelsConfig(applications)

	var panelscode string

	if utils.SourceFileExists("panels-code.html") == false {

		panelscode = CombinePanelsCode(panelcode)

		utils.WriteToSourceFile(panelscode, "panels-code.html")

	} else {
		panelscode = string(utils.ReadFromSourceFile("panels-code.html"))
	}

	utils.CopySourceToAssets("index.html", "index.html")

	utils.AddPanelsCode(panelscode)

}
