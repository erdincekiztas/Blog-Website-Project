package controllers

import (
	"goblog/client/helpers"
	"goblog/client/models"
	"goblog/services"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) HomePage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("pages/homepage")...)
	if err != nil {
		return
	}

	session, err := services.GetStore().Get(r, "system-user")
	if err != nil {
		helpers.SetAlert(w, r, "Hata!")
		http.Error(w, "/login", http.StatusSeeOther)
	}

	data := make(map[string]interface{})

	userID := session.Values["userID"]

	data["User"] = models.User{}.Get(userID)

	view.ExecuteTemplate(w, "index", data)

}
