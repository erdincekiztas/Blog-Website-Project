package controllers

import (
	"fmt"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Users struct{}

func (users Users) UserClientPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("user/users")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})

	//todo buraya danışan şeklinde ekleyeceğim şimdilik denemek için bu şekilde yapmıştım
	data["Users"] = models.User{}.GetAll(models.User{Role: "client"})

	view.ExecuteTemplate(w, "index", data)
}

/*
func (users Users) AntrenmanPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")

	var user models.User

	user = user.Get("id = ?", id)

	fmt.Println(user.Name)

}
*/
