package oauth2

import (
	"fmt"
	"html/template"

	"github.com/mmrath/gobase/common/template_util"
	"github.com/mmrath/gobase/uaa/uaa-server/internal/config"
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
	//	templatesDir := cfg.TemplateDir

	tr, err := template_util.BuildRegistry("uaa/uaa-web-react/dist")

	if err != nil {
		return nil, err
	}

	lt := tr.Get("login.html")
	if lt == nil {
		return nil, fmt.Errorf("login template not found")
	}

	ct := tr.Get("consent.html")
	if ct == nil {
		return nil, fmt.Errorf("consent template not found")
	}

	return &templateProvider{
		loginTemplate:   lt,
		consentTemplate: ct,
	}, nil

}
