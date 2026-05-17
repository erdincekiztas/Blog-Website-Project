package controllers

import (
	"encoding/json"
	"fmt"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Measurement struct{}

func (measurement Measurement) Grafik(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	userID, err := strconv.Atoi(params.ByName("userID"))
	if err != nil {
		http.Error(w, "Geçersiz kullanıcı ID", http.StatusBadRequest)
		return
	}

	// O danışana ait kullanıcı bilgisini çek
	danisan := models.User{}.Get("id = ?", userID)
	if danisan.ID == 0 {
		http.Error(w, "Danışan bulunamadı", http.StatusNotFound)
		return
	}

	view, err := template.ParseFiles(helpers.Include("user/mesurement")...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Sadece o kullanıcının ölçümlerini getir
	measurements := models.Measurement{}.GetAll("user_id = ?", userID)

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

	data := map[string]interface{}{
		"Title":        danisan.Name + " - Vücut Ölçüm Takibi",
		"UserID":       userID,
		"Labels":       template.JS(lj),
		"Weights":      template.JS(wj),
		"Heights":      template.JS(hj),
		"FatRatios":    template.JS(fj),
		"Measurements": measurements,
	}

	data["Alert"] = helpers.GetAlert(w, r)

	err = view.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (measurement Measurement) GrafikAdd(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	userID, err := strconv.Atoi(params.ByName("userID"))
	if err != nil {
		http.Error(w, "Geçersiz kullanıcı ID", http.StatusBadRequest)
		return
	}

	weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
	height, _ := strconv.ParseFloat(r.FormValue("height"), 64)
	fatRatio, _ := strconv.ParseFloat(r.FormValue("fat-ratio"), 64)

	models.Measurement{
		UserID:   uint(userID),
		Weight:   weight,
		Height:   height,
		FatRatio: fatRatio,
	}.Add()

	helpers.SetAlert(w, r, "Vücut Ölçümleri Başarıyla Eklendi!")
	http.Redirect(w, r, fmt.Sprintf("/admin/danisan/%d/grafik", userID), http.StatusSeeOther)
}
