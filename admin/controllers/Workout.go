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

type Workout struct{}

func (workout Workout) ExercisesPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	userID, err := strconv.Atoi(params.ByName("userID"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Geçersiz kullanıcı ID", http.StatusBadRequest)
		return
	}
	data := make(map[string]interface{})
	data["User"] = models.User{}.Get(uint(userID))
	data["Monday"] = models.Workout{}.GetAll(models.Workout{UserID: uint(userID), DayID: 1})
	data["Tuesday"] = models.Workout{}.GetAll(models.Workout{UserID: uint(userID), DayID: 2})
	data["Wednesday"] = models.Workout{}.GetAll(models.Workout{UserID: uint(userID), DayID: 3})
	data["Thursday"] = models.Workout{}.GetAll(models.Workout{UserID: uint(userID), DayID: 4})
	data["Friday"] = models.Workout{}.GetAll(models.Workout{UserID: uint(userID), DayID: 5})
	data["Saturday"] = models.Workout{}.GetAll(models.Workout{UserID: uint(userID), DayID: 6})
	data["Sunday"] = models.Workout{}.GetAll(models.Workout{UserID: uint(userID), DayID: 7})

	view, err := template.ParseFiles(helpers.Include("user/workout")...)

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

func (workout Workout) ExerciseAdd(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	userID, err := strconv.Atoi(params.ByName("userID"))
	if err != nil {
		http.Error(w, "Geçersiz kullanıcı ID", http.StatusBadRequest)
		return
	}

	dayID, _ := strconv.Atoi(r.FormValue("day_id"))

	models.Workout{
		UserID:     uint(userID),
		DayID:      dayID,
		Exercise:   r.FormValue("exercise"),
		Sets:       r.FormValue("sets"),
		Reps:       r.FormValue("reps"),
		RestPeriod: r.FormValue("rest_period"),
	}.Add()

	helpers.SetAlert(w, r, "Egzersiz Başarıyla Eklendi!")
	http.Redirect(w, r, fmt.Sprintf("/admin/danisan/%d/antrenman-programi", userID), http.StatusSeeOther)
}

func (workout Workout) ExerciseDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

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

	deleted_exercise := models.Workout{}.Get(uint(id))
	deleted_exercise.Delete()
	helpers.SetAlert(w, r, "Egzersiz Başarıyla Silindi!")
	http.Redirect(w, r, fmt.Sprintf("/admin/danisan/%d/antrenman-programi", userID), http.StatusSeeOther)
}
