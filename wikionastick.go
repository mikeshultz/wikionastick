/*
wikionastick

Markdown documentation Web server to serve documentation from a thumbdrive.
*/
package main

import (
	"os"
	"fmt"
	"bytes"
	"strings"
	"io/ioutil"
	"net/http"
	"path/filepath"
	log "github.com/sirupsen/logrus"
	"text/template"
	"github.com/a8m/mark"
)

var PWD string
var TEAMPLATE_DIR string

type Page struct {
	Filename string
	Title string
	Body  []byte
	HTML string
}

func loadFileToString(fname string) (string, error) {
	var bodyString string
	body, err := ioutil.ReadFile(fname)
	if err != nil {
		return "", err
	}
	bodyString = string(body)
	return bodyString, nil
}

func renderTemplate(w http.ResponseWriter, page *Page, templateName string) error {

	// Render the markdown
	page.HTML = mark.Render(string(page.Body))

	// Get full path for base template
	var templatePathBuffer bytes.Buffer

	templatePathBuffer.WriteString(PWD)
	templatePathBuffer.WriteString(TEAMPLATE_DIR)
	templatePathBuffer.WriteString("/")
	templatePathBuffer.WriteString(templateName)

	templatePath := templatePathBuffer.String()

	// Load the main template
	template_text, err := loadFileToString(templatePath)

	if err != nil {
		return err
	}
	
	// Jam it in a template
	template, err := template.New("full Thang").Parse(template_text)

	if err != nil {
		return err
	}

	err = template.Execute(w, page)

	return err

}

func renderDefaultTemplate(w http.ResponseWriter, page *Page) error {
	return renderTemplate(w, page, "base_template.html")
}

func loadPage(fname string) (*Page, error) {
	body, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return &Page{Filename: fname, Title: "CHANGEME", Body: body}, nil
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.Filename, p.Body, 0640)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	fmt.Println(string(r.URL.Path))
	var isMarkdown bool
	var path string
	path = r.URL.Path

	if path == "/" {
		// TODO: Add checks for others, like README
		path = "/index.md"
	}

	// if markdown file in not in URL, add it
	if len(r.URL.Path) > 3 && r.URL.Path[len(r.URL.Path)-3:] != ".md" {
		
		// Don't mess with CSS or JS files
		if !hasExtension(r.URL.Path, ".js") && 
			!hasExtension(r.URL.Path, ".css") {
		
			var pathBuffer bytes.Buffer
			pathBuffer.Write([]byte(r.URL.Path))
			pathBuffer.Write([]byte(".md"))
			path = pathBuffer.String()

		}

	}

	if hasExtension(path, ".md"){
		isMarkdown = true
	}
	
	fullPath := PWD + path
		
	log.WithFields(log.Fields{
		"path": fullPath,
	}).Debug("Loading file")

	if isMarkdown {

		pageOut, err := loadPage(fullPath)

		if err != nil {

			if strings.Contains(err.Error(), "no such file") {

				fmt.Fprint(w, "<h1>404</h1><div>File Not Found</div>")
				
				log.WithFields(log.Fields{
					"path": fullPath,
				}).Info("File not found")

			} else {

				fmt.Fprintf(w, "<h1>404</h1><div>File Not Found: %s</div>", err)
				
				log.WithFields(log.Fields{
					"path": fullPath,
				}).Warn("Error finding path")

			}

		} else {

			// Render markdown to html
			err = renderDefaultTemplate(w, pageOut)

			if err != nil {
			
				log.WithFields(log.Fields{
					"path": fullPath,
					"error": err,
				}).Error("Request failed")

			} else {
			
				log.WithFields(log.Fields{
					"path": fullPath,
				}).Info("Request")

			}

		}
	} else {

		// Load the file
		stringOut, err := loadFileToString(fullPath)

		if err != nil {

			fmt.Fprint(w, "<h1>404</h1><div>File Not Found</div>")
			
			log.WithFields(log.Fields{
				"path": fullPath,
			}).Info("File not found")

			return

		}

		fmt.Fprint(w, stringOut)

	}
}

func init() {

	// Get pwd
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	PWD = filepath.Dir(ex)					// Current working directory
	TEAMPLATE_DIR = "/templates/default/"	// Default template directory

	// Setup logging
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

}

func main() {
	log.Info("Server init")
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8888", nil)
}