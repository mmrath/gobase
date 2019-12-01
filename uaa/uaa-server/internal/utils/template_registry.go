package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/oxtoacart/bpool"
)

type templateRegistry struct {
	templates  map[string]*template.Template
	bufferPool *bpool.BufferPool
}

func NewTemplateRegistry() (*templateRegistry, error) {
	templates := make(map[string]*template.Template)
	templateDir := "uaa/uaa-server/resources/web/templates"

	layoutDir := filepath.Join(templateDir, "layout")

	layoutFiles, err := filepath.Glob(filepath.Join(layoutDir, "*.html"))
	if err != nil {
		return nil, err
	}

	layoutTemplates := template.Must(template.ParseFiles(layoutFiles...))

	layoutClone := template.Must(layoutTemplates.Clone())

	tmplName := "account/sign-up-form.html"
	tmplPath := filepath.Join(templateDir, tmplName)
	templates[tmplName] = template.Must(layoutClone.ParseFiles(tmplPath))

	return &templateRegistry{
		templates:  templates,
		bufferPool: bpool.NewBufferPool(64),
	}, nil
}




func (r *templateRegistry) RenderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	templateDir := "uaa/uaa-web-app/build/"
	tmplPath := filepath.Join(templateDir, name)
	tmpl := template.Must(template.ParseFiles(tmplPath))


	buf := r.bufferPool.Get()
	defer r.bufferPool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		err := fmt.Errorf("template execution failed")
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
	return nil
}

