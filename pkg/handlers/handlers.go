package handlers

import (
	"net/http"

	"github.com/iam-vl/go-web/pkg/config"
	"github.com/iam-vl/go-web/pkg/models"
	"github.com/iam-vl/go-web/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NH sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Business logic here
	strMap := make(map[string]string)
	strMap["test1"] = "Hello again"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	strMap["remote_ip"] = remoteIp

	// Send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: strMap,
	})
}
