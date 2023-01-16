// Модуль agent отправляет информацию о состоянии
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/vllvll/keepa/internal/pages"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

const BuildTemplate = `
Build version: {{ .version }}
Build date: {{ .date }}
Build commit: {{ .commit }}
`

func main() {
	t := template.Must(template.New("build").Parse(BuildTemplate))
	err := t.Execute(os.Stdout, map[string]string{
		"version": buildVersion,
		"date":    buildDate,
		"commit":  buildCommit,
	})
	if err != nil {
		log.Fatalf("Error with config: %v", err)
	}

	authModel := pages.NewAuthModel()
	p := tea.NewProgram(authModel)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
