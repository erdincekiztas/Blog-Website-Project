package controllers

import (
	"crypto/sha256"
	"fmt"
	"goblog/admin/models"
	"goblog/client/helpers"
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

	if user.Password == password && user.Role == "client" {

		services.SetUser(w, r, user.ID, user.Email, user.Role)

		message := "Hoşgeldiniz " + user.Name
		helpers.SetAlert(w, r, message)
		http.Redirect(w, r, "/anasayfa", http.StatusSeeOther)

	} else {
		helpers.SetAlert(w, r, "Kullanıcı Adı ya da Şifre Yanlış")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

func (authorization Authorization) RegistrationPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("authorization/registration")...)
	if err != nil {
		http.Error(w, "sayfa yüklenemedi!", http.StatusInternalServerError)
	}

	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)

	view.ExecuteTemplate(w, "index", data)

}
func (authorization Authorization) Registration(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := strings.TrimSpace(r.FormValue("name"))
	mail := strings.TrimSpace(r.FormValue("mail"))
	password := r.FormValue("password")
	password_confirm := r.FormValue("password_confirm")

	if password != password_confirm {
		helpers.SetAlert(w, r, "Şifreler Eşleşmiyor!")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}
	user := models.User{}.Get(models.User{Email: mail})
	if user.Email != "" {
		helpers.SetAlert(w, r, "Bu mail adresi zaten kullanılıyor!")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}
	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	// Yeni kullanıcı oluştur
	newUser := models.User{
		Name:     name,
		Email:    mail,
		Password: hashedPassword,
	}
	newUser.Add()

	helpers.SetAlert(w, r, "Kayıt başarılı! Giriş yapabilirsiniz.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (authorization Authorization) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	services.RemoveUser(w, r)
	helpers.SetAlert(w, r, "Çıkış Yapıldı")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
