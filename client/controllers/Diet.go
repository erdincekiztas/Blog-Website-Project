package controllers

import (
	"goblog/client/helpers"
	"goblog/client/models"
	"goblog/services"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Diet struct{}

func (diet Diet) DietPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("pages/diet")...)
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

	user := models.User{}.Get(userID)
	data["User"] = user

	data["Breakfast"] = models.Diet{}.GetAll(models.Diet{UserID: user.ID, MealID: 1})
	data["Snack"] = models.Diet{}.GetAll(models.Diet{UserID: user.ID, MealID: 2})
	data["Lunch"] = models.Diet{}.GetAll(models.Diet{UserID: user.ID, MealID: 3})
	data["Dinner"] = models.Diet{}.GetAll(models.Diet{UserID: user.ID, MealID: 4})
	data["NightMeal"] = models.Diet{}.GetAll(models.Diet{UserID: user.ID, MealID: 5})
	data["PreWorkout"] = models.Diet{}.GetAll(models.Diet{UserID: user.ID, MealID: 6})
	data["PostWorkout"] = models.Diet{}.GetAll(models.Diet{UserID: user.ID, MealID: 7})

	view.ExecuteTemplate(w, "index", data)

}
