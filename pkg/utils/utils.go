package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GenerateMainPage(HtmlText string) {

	filePath := sourcePathPrefix + "index.html"

	f, err := os.Create(filePath)

	if err != nil {
		panic(fmt.Errorf("Cannot create index.html - %w", err))
	}

	defer f.Close()

	n3, err := f.WriteString(HtmlText)

	if err != nil {
		panic(fmt.Errorf("Cannot write to index.html - %w", err))
	}

	log.Printf("wrote %d bytes\n", n3)

	f.Close()

	log.Println("index.html file main code generated")

	return

}

func AddPanelsCode(HtmlText string) {

	f, err := os.OpenFile("assets/index.html", os.O_RDWR, 0644)

	log.Println("index.html opened to add panels code")

	if err != nil {
		panic(fmt.Errorf("Cannot open index.html - %w", err))
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	var buffer strings.Builder

	for scanner.Scan() {

		line := scanner.Text()

		buffer.WriteString(line + "\n")

		if strings.Contains(line, "PANELS") {

			buffer.WriteString(HtmlText)
		}
	}

	f.Close()
	err = os.WriteFile("assets/index.html", []byte(buffer.String()), 0644)

	if err != nil {
		panic(fmt.Errorf("Cannot update index.html - %w", err))
	}

	f.Close()

	log.Println("panels code generated")

	return
}
