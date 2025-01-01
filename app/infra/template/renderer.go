package template

import (
	"ipr/shared"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views *jet.Set

func InitializeRenderer(templateDir string, devMode bool) {
	if devMode {
		// Development mode: reloads templates automatically
		views = jet.NewSet(
			jet.NewOSFileSystemLoader(templateDir),
			jet.InDevelopmentMode(),
		)
	} else {
		// Production mode: optimized for performance
		views = jet.NewSet(
			jet.NewOSFileSystemLoader(templateDir),
		)
	}
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	data["asset_hash"] = func(asset string) string {
		return asset + "?v=12345" //todo
	}

	tmpl, err := views.GetTemplate(templateName)
	if err != nil {
		shared.HandleHttpError(w, r, err, nil)

		return
	}

	if err := tmpl.Execute(w, nil, data); err != nil {
		shared.HandleHttpError(w, r, err, nil)
	}
}
