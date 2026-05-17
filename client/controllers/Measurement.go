package controllers

import (
	"encoding/json"
	"fmt"
	"goblog/client/helpers"
	"goblog/client/models"
	"goblog/services"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Measurement struct{}

func (measurement Measurement) MeasurementPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("pages/measurement")...)
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

	measurements := models.Measurement{}.GetAll("user_id = ?", user.ID)

	var labels, weights, heights, fatRatios []string
	for _, m := range measurements {
		labels = append(labels, m.CreatedAt.Format("02 Jan 2006"))
		weights = append(weights, fmt.Sprintf("%.1f", m.Weight))
		heights = append(heights, fmt.Sprintf("%.1f", m.Height))
		fatRatios = append(fatRatios, fmt.Sprintf("%.1f", m.FatRatio))
	}

	lj, _ := json.Marshal(labels)
	wj, _ := json.Marshal(weights)
	hj, _ := json.Marshal(heights)
	fj, _ := json.Marshal(fatRatios)

	data["Title"] = " Vücut Ölçümleriniz"
	data["User"] = user
	data["Labels"] = template.JS(lj)
	data["Weights"] = template.JS(wj)
	data["Heights"] = template.JS(hj)
	data["FatRatios"] = template.JS(fj)
	data["Measurements"] = measurements

	view.ExecuteTemplate(w, "index", data)

}
