package oauth2

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/mmrath/gobase/common/template_util"
	"github.com/mmrath/gobase/uaa/uaa-server/internal/config"
)

var templateDir string = "uaa/uaa-web-app/build/"


type templateProvider struct {
	loginTemplate   *template.Template
	consentTemplate *template.Template
}

func (t *templateProvider) LoginTemplate() *template.Template {
	tmplPath := filepath.Join(templateDir, "oauth2/login.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))

	return tmpl
}

func (t *templateProvider) ConsentTemplate() *template.Template {
	tmplPath := filepath.Join(templateDir, "oauth2/consent.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	return tmpl
}

func loadTemplates(cfg config.WebConfig) (TemplateProvider, error) {
	//	templatesDir := cfg.TemplateDir
	templateDir = cfg.TemplateDir
	tr, err := template_util.BuildRegistry(cfg.TemplateDir)

	if err != nil {
		return nil, err
	}

	lt := tr.Get("oauth2/login.html")
	if lt == nil {
		return nil, fmt.Errorf("login template not found")
	}

	ct := tr.Get("oauth2/consent.html")
	if ct == nil {
		return nil, fmt.Errorf("consent template not found")
	}

	return &templateProvider{
		loginTemplate:   lt,
		consentTemplate: ct,
	}, nil

}
