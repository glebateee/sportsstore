package admin

import (
	"platform/http/actionresults"
	"platform/http/handling"
)

var sectionNames = []string{"Products", "Categories", "Orders", "Database"}

type AdminHandler struct {
	handling.URLGenerator
}

type AdminTemplateContext struct {
	Sections       []string
	ActiveSection  string
	SectionUrlFunc func(string) string
}

func (handler AdminHandler) GetSection(section string) actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin.html", AdminTemplateContext{
		Sections:      sectionNames,
		ActiveSection: section,
		SectionUrlFunc: func(tgt string) (url string) {
			url, _ = handler.URLGenerator.GenerateUrl(AdminHandler.GetSection, tgt)
			return
		},
	})
}
