package utils

import (
	"bytes"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/nicksnyder/go-i18n/i18n"
	"html/template"
	"net/http"
	"path/filepath"
)

// Global storage for templates
var htmlTemplates *template.Template
var htmlTemplatePath string
var htmlTemplateDirName string = "templates"

type HTMLTemplate struct {
	TemplateName string
	Props        map[string]interface{}
	Html         map[string]template.HTML
	Locale       string
}

func InitHTML() {
	templatesDir := FindDir(htmlTemplateDirName)
	htmlTemplatePath, _ = filepath.Abs(templatesDir)
	htmlTemplatePath += "/"
	l4g.Debug(T("api.api.init.parsing_templates.debug"), htmlTemplatePath)
	initHTMLWithDir(htmlTemplatePath)
}

func dirChangeNotify(dir string) {
	if dir != htmlTemplatePath {
		l4g.Error(fmt.Sprintf("file watcher has something error"))
		return
	}
	dir += "/"
	var err error
	if htmlTemplates, err = template.ParseGlob(dir + "*.html"); err != nil {
		l4g.Error(T("web.parsing_templates.error"), err)
	}
}

func initHTMLWithDir(dir string) {
	if htmlTemplates != nil {
		return
	}
	var err error
	if htmlTemplates, err = template.ParseGlob(dir + "*.html"); err != nil {
		l4g.Error(T("api.api.init.parsing_templates.error"), err)
	}
}

func NewHTMLTemplate(templateName string, locale string) *HTMLTemplate {
	return &HTMLTemplate{
		TemplateName: templateName,
		Props:        make(map[string]interface{}),
		Html:         make(map[string]template.HTML),
		Locale:       locale,
	}
}

func (t *HTMLTemplate) addDefaultProps() {
	var localT i18n.TranslateFunc
	if len(t.Locale) > 0 {
		localT = GetUserTranslations(t.Locale)
	} else {
		localT = T
	}
	t.Props["Footer"] = localT("api.templates.email_footer")
	t.Props["Organization"] = localT("api.templates.email_organization")
	t.Html["EmailInfo"] = template.HTML(localT("api.templates.email_info",
		map[string]interface{}{"SupportEmail": "primefour@163.com", "SiteName": "www.fpbbc.com"}))
}

func (t *HTMLTemplate) Render() string {
	t.addDefaultProps()

	var text bytes.Buffer

	if err := htmlTemplates.ExecuteTemplate(&text, t.TemplateName, t); err != nil {
		l4g.Error(T("api.api.render.error"), t.TemplateName, err)
	}

	return text.String()
}

func (t *HTMLTemplate) RenderToWriter(w http.ResponseWriter) error {
	t.addDefaultProps()

	if err := htmlTemplates.ExecuteTemplate(w, t.TemplateName, t); err != nil {
		l4g.Error(T("api.api.render.error"), t.TemplateName, err)
		return err
	}
	return nil
}
