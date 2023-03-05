package handlers

import (
	"net/http"
	"web3/models"
	"web3/pkg/config"
	"web3/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		App: appConfig,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, request *http.Request) {
	m.App.Session.Put(request.Context(), "userid", "marcinemski")

	render.RenderTemplate(w, "home.page.html", &models.PageData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter, request *http.Request) {

	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves. We love it!"

	userid := m.App.Session.GetString(request.Context(), "userid")
	strMap["userid"] = userid

	render.RenderTemplate(w, "about.page.html", &models.PageData{StrMap: strMap})
}
