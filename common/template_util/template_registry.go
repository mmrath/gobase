package template_util

import (
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/common/error_util"
)

type Registry struct {
	templates *template.Template
}

func (t *Registry) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (t *Registry) Get(name string) *template.Template {
	return t.templates.Lookup(name)
}

func BuildRegistry(templateDir string) (*Registry, error) {
	tmpl, err := findAndParseTemplates(templateDir, template.FuncMap{})

	if err != nil {
		return nil, error_util.NewInternal(err, "failed to load template")
	}
	return &Registry{templates: tmpl}, nil
}

func findAndParseTemplates(rootDir string, funcMap template.FuncMap) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
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

func (t *Registry) RenderHttp(w http.ResponseWriter, templateName string, data interface{}) {
	err := t.Render(w, templateName, data)
	if err != nil {
		err = error_util.NewInternal(err, "failed to render template")
		log.Error().
			Err(err).
			Str("template", templateName).
			Interface("data", data).
			Msg("failed to render template")
		w.WriteHeader(http.StatusInternalServerError)
		err = t.Render(w, "error/500.html", error_util.GetErrorID(err))

	}
}
