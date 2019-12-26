package template_util

import (
	"github.com/unrolled/render"
	"io"
)

type templateRegistry struct {
	render *render.Render
}
type AssetFunc func(string) ([]byte, error)
type AssetNamesFunc func() []string

func NewTemplateRegistry(assetFunc AssetFunc, assetNamesFunc AssetNamesFunc, layout string) *templateRegistry {
	r := render.New(render.Options{
		Asset:      assetFunc,
		AssetNames: assetNamesFunc,
		Extensions: []string{".tmpl", ".html"},
		Layout:     layout,
	})
	return &templateRegistry{r}
}

func (t *templateRegistry) RenderHTML(w io.Writer, status int, name string, binding interface{}) error {
	return t.render.HTML(w, status, name, binding)
}
