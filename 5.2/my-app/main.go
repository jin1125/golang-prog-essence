package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl := `{{.}}`
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, "Hello, World!")
	if err != nil {
		log.Fatal(err)
	}
}
