package handlers

import (
	"net/http"
	"web3/models"
	render "web3/pkg/render"
)

func HomeHandler(w http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.PageData{})
}

func AboutHandler(w http.ResponseWriter, request *http.Request) {

	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves. We love it!"

	render.RenderTemplate(w, "about.page.html", &models.PageData{StrMap: strMap})
}
