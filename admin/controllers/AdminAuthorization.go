package controllers

import (
	"crypto/sha256"
	"fmt"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"goblog/services"
	"html/template"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Authorization struct{}

func (authorization Authorization) LoginPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("authorization/login")...)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "sayfa yüklenemedi!", http.StatusInternalServerError)
	}
	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)

	view.ExecuteTemplate(w, "index", data)

}

func (authorization Authorization) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	mail := strings.TrimSpace(r.FormValue("mail"))
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))

	user := models.User{}.Get(models.User{
		Email: mail,
	})
	fmt.Println(user)
	if user.Password == password && user.Role == "admin" {

		services.SetUser(w, r, user.ID, user.Email, user.Role)

		message := "Hoşgeldiniz " + user.Name
		helpers.SetAlert(w, r, message)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)

	} else {
		helpers.SetAlert(w, r, "Kullanıcı Adı ya da Şifre Yanlış")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}

}

func (authorization Authorization) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	services.RemoveUser(w, r)
	helpers.SetAlert(w, r, "Çıkış Yapıldı")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}
