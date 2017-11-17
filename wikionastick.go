/*
wikionastick

Markdown documentation Web server to serve documentation from a thumbdrive.
*/
package main

import (
	"os"
	"path"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"text/template"
	"github.com/a8m/mark"
)

var PWD string
var TEMPLATE_DIR string

type Page struct {
	Filename string
	Title string
	Body  []byte
	HTML string
}

func loadFileToString(fullPath string) (string, error) {
	var bodyString string
	body, err := ioutil.ReadFile(fullPath)
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
	templatePath := path.Join(PWD, TEMPLATE_DIR, templateName)

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
	return &Page{Filename: fname, Title: "WikiOnAStick", Body: body}, nil
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.Filename, p.Body, 0640)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	fmt.Println(string(r.URL.Path))
	var isMarkdown bool
	var filePath string
	filePath = r.URL.Path

	// if markdown file in not in URL, add it
	if len(filePath) > 3 && filePath[len(filePath)-3:] != ".md" {
		
		// Don't mess with CSS or JS files
		if !HasExtension(r.URL.Path, ".js") && 
			!HasExtension(r.URL.Path, ".css") {

			filePath = path.Join(r.URL.Path, ".md")

		}

	}
	
	var fullPath string
	
	// Handle default page
	if filePath == "/" {

		// First try is index.md
		filePath = "/index.md"

		// Check if index.md is a file
		if _, err := os.Stat(PWD + filePath); !os.IsNotExist(err) {
			fullPath = PWD + filePath
		} else {
			
			// Let's fallback to README.md
			filePath = "/README.md"

			// Check if README.md exists
			if _, err := os.Stat(PWD + filePath); !os.IsNotExist(err) {
				fullPath = PWD + filePath
			} else {
				// All Else Failed is an awesome band
				http.Error(w, "File not found", 404)
			}

		}
	} else {
		fullPath = PWD + filePath
	}

	if HasExtension(filePath, ".md"){
		isMarkdown = true
	}
		
	log.WithFields(log.Fields{
		"filePath": fullPath,
	}).Debug("Loading file")

	if isMarkdown {

		pageOut, err := loadPage(fullPath)

		if err != nil {

			if strings.Contains(err.Error(), "no such file") {

				http.Error(w, "File not found", 404)
				
				log.WithFields(log.Fields{
					"fullPath": fullPath,
				}).Info("File not found")

			} else {

				http.Error(w, "File not found", 404)
				
				log.WithFields(log.Fields{
					"fullPath": fullPath,
				}).Warn("Error finding path")

			}

		} else {

			// Render markdown to html
			err = renderDefaultTemplate(w, pageOut)

			if err != nil {
			
				log.WithFields(log.Fields{
					"fullPath": fullPath,
					"error": err,
				}).Error("Request failed")

			} else {
			
				log.WithFields(log.Fields{
					"fullPath": fullPath,
				}).Info("Request")

			}

		}
	} else {

		// Load the file
		stringOut, err := loadFileToString(fullPath)

		if err != nil {

			http.Error(w, "File not found", 404)
			
			log.WithFields(log.Fields{
				"fullPath": fullPath,
			}).Info("File not found")

			return

		}

		fmt.Fprint(w, stringOut)

	}
}

func init() {

	var opts struct {
		TemplateDir string `short:"t" long:"template" description:"Path to a template directory"`
		LogLevel string `short:"l" long:"loglevel" description:"Log level. ['debug', 'info', 'warning', 'error']"`
	}

	// Get pwd
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	PWD = filepath.Dir(ex)					// Current working directory

	args, err := flags.Parse(&opts)

	if err != nil {

		log.WithFields(log.Fields{
			"args": args,
			"sysargs": os.Args[1:],
		}).Error("Error parsing args")

		os.Exit(1)
	}

	TEMPLATE_DIR = opts.TemplateDir
	logLevel := opts.LogLevel

	// Setup logging
	log.SetOutput(os.Stdout)
	log.SetLevel(LogLevelTranslate(logLevel))

	log.WithFields(log.Fields{
		"args": args,
		"sysargs": os.Args[1:],
	}).Debug("Args")

}

func main() {
	log.Info("Server init")
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8888", nil)
}