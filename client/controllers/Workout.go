package controllers

import (
	"fmt"
	"goblog/client/helpers"
	"goblog/client/models"
	"goblog/services"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Workout struct{}

func (workout Workout) ExercisesPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	data := make(map[string]interface{})

	session, err := services.GetStore().Get(r, "system-user")
	if err != nil {
		helpers.SetAlert(w, r, "Hata!")
		http.Error(w, "/login", http.StatusSeeOther)
	}

	userID := session.Values["userID"]
	user := models.User{}.Get(userID)

	data["User"] = user

	data["Monday"] = models.Workout{}.GetAll(models.Workout{UserID: user.ID, DayID: 1})
	data["Tuesday"] = models.Workout{}.GetAll(models.Workout{UserID: user.ID, DayID: 2})
	data["Wednesday"] = models.Workout{}.GetAll(models.Workout{UserID: user.ID, DayID: 3})
	data["Thursday"] = models.Workout{}.GetAll(models.Workout{UserID: user.ID, DayID: 4})
	data["Friday"] = models.Workout{}.GetAll(models.Workout{UserID: user.ID, DayID: 5})
	data["Saturday"] = models.Workout{}.GetAll(models.Workout{UserID: user.ID, DayID: 6})
	data["Sunday"] = models.Workout{}.GetAll(models.Workout{UserID: user.ID, DayID: 7})

	view, err := template.ParseFiles(helpers.Include("pages/workout")...)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Sayfa yüklenemedi", http.StatusInternalServerError)
		return
	}

	view.ExecuteTemplate(w, "index", data)
}
