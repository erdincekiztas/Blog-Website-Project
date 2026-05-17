package controllers

import (
	"fmt"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Diet struct{}

func (diet Diet) DietPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	userID, err := strconv.Atoi(params.ByName("userID"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Geçersiz kullanıcı ID", http.StatusBadRequest)
		return
	}

	data := make(map[string]interface{})

	data["User"] = models.User{}.Get(uint(userID))

	data["Breakfast"] = models.Diet{}.GetAll(models.Diet{UserID: uint(userID), MealID: 1})
	data["Snack"] = models.Diet{}.GetAll(models.Diet{UserID: uint(userID), MealID: 2})
	data["Lunch"] = models.Diet{}.GetAll(models.Diet{UserID: uint(userID), MealID: 3})
	data["Dinner"] = models.Diet{}.GetAll(models.Diet{UserID: uint(userID), MealID: 4})
	data["NightMeal"] = models.Diet{}.GetAll(models.Diet{UserID: uint(userID), MealID: 5})
	data["PreWorkout"] = models.Diet{}.GetAll(models.Diet{UserID: uint(userID), MealID: 6})
	data["PostWorkout"] = models.Diet{}.GetAll(models.Diet{UserID: uint(userID), MealID: 7})

	view, err := template.ParseFiles(helpers.Include("user/diet")...)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Sayfa yüklenemedi", http.StatusInternalServerError)
		return
	}

	data["Alert"] = helpers.GetAlert(w, r)

	if err := view.ExecuteTemplate(w, "index", data); err != nil {
		fmt.Println(err)
	}

}

func (diet Diet) DietAdd(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	userID, err := strconv.Atoi(params.ByName("userID"))

	if err != nil {
		http.Error(w, "Geçersiz kullanıcı ID", http.StatusBadRequest)
		return
	}

	mealID, _ := strconv.Atoi(r.FormValue("meal_id"))

	models.Diet{
		UserID:   uint(userID),
		MealID:   mealID,
		Food:     r.FormValue("food"),
		Grammage: r.FormValue("grammage"),
	}.Add()
	helpers.SetAlert(w, r, "Öğün Başarıyla Eklendi!")
	http.Redirect(w, r, fmt.Sprintf("/admin/danisan/%d/diyet-programi", userID), http.StatusSeeOther)
}

func (diet Diet) DietDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	userID, err := strconv.Atoi(params.ByName("userID"))

	if err != nil {
		http.Error(w, "Geçersiz kullanıcı ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		http.Error(w, "Geçersiz ID", http.StatusBadRequest)
		return
	}

	deletedDiet := models.Diet{}.Get(uint(id))
	deletedDiet.Delete()
	helpers.SetAlert(w, r, "Öğün Başarıyla Silindi!")
	http.Redirect(w, r, fmt.Sprintf("/admin/danisan/%d/diyet-programi", userID), http.StatusSeeOther)
}
