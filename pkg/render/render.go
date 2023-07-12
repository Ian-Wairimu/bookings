package render

import (
	"fmt"
	"github.com/Ian-Wairimu/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
)

//var app *config.AppConfig
//
////NewTemplates sets the config for the template package
//
//func NewTemplates(a *config.AppConfig) {
//	app = a
//}

// TemplateTmplTest is used to render Templates in golang
//func TemplateTmplTest(w http.ResponseWriter, tmpl string) {
//	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
//	err := parsedTemplate.Execute(w, nil)
//	if err != nil {
//		log.Fatal("error parsing template")
//	}
//}

var tc = make(map[string]*template.Template)

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
func TemplateTmpl(w http.ResponseWriter, t string, td *models.TemplateData) {
	//want to get the template cache from the AppConfig
	var tmpl *template.Template
	var err error

	//  want to check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		//need to create the template
		log.Println("creating template and adding to cache")
		err, _ = CreateTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}
	tmpl = tc[t]
	td = AddDefaultData(td)
	err = tmpl.Execute(w, td)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache the template cache
func CreateTemplateCache(t string) (error, error) {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err, nil
	}
	// add template to cache
	tc[t] = tmpl
	return nil, nil
}
