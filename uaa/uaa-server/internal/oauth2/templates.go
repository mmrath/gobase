package oauth2

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"

	"github.com/mmrath/gobase/uaa-server/internal/config"
)

type templateProvider struct {
	loginTemplate   *template.Template
	consentTemplate *template.Template
}

func (t *templateProvider) LoginTemplate() *template.Template {
	return t.loginTemplate
}

func (t *templateProvider) ConsentTemplate() *template.Template {
	return t.consentTemplate
}

func loadTemplates(cfg config.WebConfig) (TemplateProvider, error) {
	templatesDir := cfg.TemplateDir

	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read tempalte dir %s: %w", templatesDir, err)
	}

	filenames := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filenames = append(filenames, filepath.Join(templatesDir, file.Name()))
	}
	if len(filenames) == 0 {
		return nil, fmt.Errorf("no files in template dir %q", templatesDir)
	}

	tmpls, err := template.New("").ParseFiles(filenames...)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template files: %w", err)
	}

	loginTmpl := tmpls.Lookup("login.html")
	consentTmpl := tmpls.Lookup("consent.html")

	return &templateProvider{loginTemplate: loginTmpl, consentTemplate: consentTmpl}, nil

}
