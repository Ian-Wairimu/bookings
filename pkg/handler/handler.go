package handler

import (
	"github.com/Ian-Wairimu/bookings/pkg/config"
	"github.com/Ian-Wairimu/bookings/pkg/models"
	"github.com/Ian-Wairimu/bookings/pkg/render"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers it set's the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//here we render the homepage template
	render.TemplateTmpl(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perfoming some logic
	stringMap := make(map[string]string)
	remoteIp := r.RemoteAddr
	//m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	stringMap["test"] = remoteIp
	render.TemplateTmpl(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
