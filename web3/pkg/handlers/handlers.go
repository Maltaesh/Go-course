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
	render.RenderTemplate(w, "home.page.html", &models.PageData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter, request *http.Request) {

	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves. We love it!"

	render.RenderTemplate(w, "about.page.html", &models.PageData{StrMap: strMap})
}
