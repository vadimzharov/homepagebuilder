package pagebuilder

import (
	"bytes"
	"html/template"
	"log"
	"strings"

	//"text/template"

	_ "encoding/json"

	"github.com/vadimzharov/homepagebuilder/pkg/config"
)

func panelTemplateValid(panelcode string) bool {

	DBConfig, _ = config.ReadConfig()

	tmpl, err := template.New("test").Parse(panelcode)
	if err != nil {
		return false
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, DBConfig.Panels[1]); err != nil {
		return false
	}

	return true
}

func GeneratePanelsConfig(apps []Applications) {
	DBConfig, _ = config.ReadConfig()

	for i := range DBConfig.Panels {

		log.Println("Creating config for panel named", DBConfig.Panels[i].Name)
		for j := range apps {
			if DBConfig.Panels[i].Name == apps[j].Name {
				DBConfig.Panels[i].Description = apps[j].Description
			}
		}

	}

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
