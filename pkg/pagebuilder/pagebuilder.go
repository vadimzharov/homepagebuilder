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

var mainpagequery string
var panelquery string

func BuildNewPage() {

	dbconfig, customQueries := config.ReadConfig()

	if customQueries.Mainpagequery != "" {
		mainpagequery = customQueries.Mainpagequery
	} else {
		mainpagequery = Defaultmainpagequery
	}

	if customQueries.Panelquery != "" {
		panelquery = customQueries.Panelquery
	} else {
		panelquery = Defaultpanelquery
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

		}

		appsdescription = []byte(chatgpt.QueryGPT(dbconfig.ChatGPTAPIKey, appsquery.String()))

		utils.WriteToSourceFile(string(appsdescription), "appsdescription.json")

	} else {
		appsdescription = utils.ReadFromSourceFile("appsdescription.json")
	}

	err := json.Unmarshal([]byte(appsdescription), &applications)
	if err != nil {
		fmt.Println(err)
		return
	}

	var panelcode string

	if utils.SourceFileExists("panel-code.tmpl") == false {

		panelcode = chatgpt.QueryGPT(dbconfig.ChatGPTAPIKey, panelquery)

		c := 0

		ispaneltemplatevalid := panelTemplateValid(panelcode)

		for !ispaneltemplatevalid {
			if c > 3 {
				log.Panic("Cannot validate panel's Golang template generated by ChatGTP even after quering it for ", c, "times. probably there are variables present not supported by configuration. The code is below:", panelcode)
				break
			}

			panelcode = chatgpt.QueryGPT(dbconfig.ChatGPTAPIKey, panelquery)

			ispaneltemplatevalid = panelTemplateValid(panelcode)
			c += 1

		}

		utils.WriteToSourceFile(panelcode, "panel-code.tmpl")

	} else {

		log.Println("Reading panel code from already generated template...")

		panelcode = string(utils.ReadFromSourceFile("panel-code.tmpl"))

		c := 0

		ispaneltemplatevalid := panelTemplateValid(panelcode)

		for !ispaneltemplatevalid {
			if c > 3 {
				log.Panic("Cannot validate panel's Golang template generated by ChatGTP even after quering it for ", c, "times. probably there are variables present not supported by configuration. The code is below:", panelcode)
				break
			}

			panelcode = chatgpt.QueryGPT(dbconfig.ChatGPTAPIKey, panelquery)

			ispaneltemplatevalid = panelTemplateValid(panelcode)
			c += 1

		}

		utils.WriteToSourceFile(panelcode, "panel-code.tmpl")
	}

	GeneratePanelsConfig(applications)

	var panelscode string

	if utils.SourceFileExists("panels-code.html") == false {

		panelscode = CombinePanelsCode(panelcode)

		utils.WriteToSourceFile(panelscode, "panels-code.html")

	} else {
		panelscode = string(utils.ReadFromSourceFile("panels-code.html"))
	}

	if utils.SourceFileExists("index.html") == false {

		landingpagecode := chatgpt.QueryGPT(dbconfig.ChatGPTAPIKey, mainpagequery)

		utils.WriteToSourceFile(landingpagecode, "index.html")
		utils.CopySourceToAssets("index.html", "index.html")

	} else {
		utils.CopySourceToAssets("index.html", "index.html")
	}

	utils.AddPanelsCode(panelscode)

}
