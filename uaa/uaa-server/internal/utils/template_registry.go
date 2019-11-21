package utils

import (
	"fmt"
	"github.com/oxtoacart/bpool"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var templates map[string]*template.Template
var bufpool *bpool.BufferPool
var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

// create a buffer pool
func init() {
	bufpool = bpool.NewBufferPool(64)
	log.Println("buffer allocation successful")
}

type TemplateConfig struct {
	TemplateLayoutPath  string
	TemplateIncludePath string
}

type TemplateError struct {
	s string
}

func (e *TemplateError) Error() string {
	return e.s
}

func NewError(text string) error {
	return &TemplateError{text}
}

var templateConfig *TemplateConfig

func SetTemplateConfig(layoutPath, includePath string) {
	templateConfig = &TemplateConfig{layoutPath, includePath}
}

func LoadTemplates() (err error) {

	if templateConfig == nil {
		err = NewError("TemplateConfig not initialized")
		return err
	}
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(templateConfig.TemplateLayoutPath + "*.html")
	if err != nil {
		return err
	}

	includeFiles, err := filepath.Glob(templateConfig.TemplateIncludePath + "*.html")
	if err != nil {
		return err
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(mainTmpl)

	if err != nil {
		log.Fatal(err)
	}

	layoutTemplates := template.Must(template.ParseFiles(layoutFiles...))

	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			return err
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}

	log.Println("templates loading successful")
	return nil

}


func findAndParseTemplates(layoutPath, includePath string, funcMap template.FuncMap) error {
	cleanRoot := filepath.Clean(includePath)
	pfx := len(cleanRoot) + 1

	layoutFiles, err := filepath.Glob(layoutPath + "/*.html")

	if err != nil {
		return err
	}
	layoutTemplates := template.Must(template.ParseFiles(layoutFiles...))
	root, err := layoutTemplates.Clone()

	if err != nil {
		return err
	}

	err = filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t, err := root.Clone()

			if err != nil {
				return err
			}

			t := root.New(name).Funcs(funcMap)


			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}
		return nil
	})

	return root, err
}


func RenderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
		err := NewError("Template doesn't exist")
		return err
	}

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		err := NewError("Template execution failed")
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
	return nil
}