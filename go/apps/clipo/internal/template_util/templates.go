package template_util

import (
	"bytes"
	"github.com/mmrath/gobase/go/apps/clipo/internal/generated"
	"github.com/mmrath/gobase/go/pkg/errutil"
	"html/template"
	"io"
)

var tmpl = make(map[string]*template.Template)

func init() {
	tmpl["index.html"] = template.Must(template.ParseFiles("index.html", "base.html"))
	tmpl["other.html"] = template.Must(template.ParseFiles("other.html", "base.html"))
}

type Registry struct {
	templateMap map[string]*template.Template
}

func NewRegistry() (*Registry, error) {
	templateMap := make(map[string]*template.Template)
	layoutFile := "templates/email/layout.gohtml"

	layoutTmpl, err := generated.Asset(layoutFile)
	if err != nil {
		return nil, errutil.Wrapf(err, "unable to load %s", layoutFile)
	}

	tmpl, err := template.New("").Parse(string(layoutTmpl))

	if err != nil {
		return nil, errutil.Wrapf(err, "failed to parse %s", layoutFile)
	}

	file := "templates/email/auth/account_activation.gohtml"
	fileData, err := generated.Asset(file)
	if err != nil {
		return nil, errutil.Wrapf(err, "unable to load %s", file)
	}

	templateMap[file], err = tmpl.Parse(string(fileData))

	if err != nil {
		return nil, errutil.Wrapf(err, "failed to parse %s", file)
	}
	return &Registry{templateMap: templateMap}, nil
}

func (t *Registry) Render(w io.Writer, name string, data interface{}) error {
	tmpl := t.templateMap[name]

	if tmpl == nil {
		return errutil.Errorf("failed to render template %s", name)
	}

	err := tmpl.Execute(w, data)

	if err != nil {
		return errutil.Wrapf(err, "failed to render template %s", name)
	}
}

func (t *Registry) RenderToString(name string, data interface{}) (string, error) {
	buf := new(bytes.Buffer)

	err := t.Render(buf, name, data)

	if err != nil {
		return "", errutil.Wrapf(err, "failed to render template to string %s", name)
	}
}
