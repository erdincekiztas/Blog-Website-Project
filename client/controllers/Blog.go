package controllers

import (
	"fmt"
	"goblog/client/helpers"
	"goblog/client/models"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Blog struct{}

func (blog Blog) BlogPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("blogsite/list")...)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Sayfa yüklenirken bir hata oluştu!", http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})

	data["Posts"] = models.Post{}.GetAll()

	view.ExecuteTemplate(w, "index", data)

}

func (blog Blog) Detail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("blogsite/detail")...)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Sayfa yüklenirken bir hata oluştu!", http.StatusInternalServerError)
		return
	}
	slug := params.ByName("slug")

	data := make(map[string]interface{})
	post := models.Post{}.Get("slug = ?", slug)
	data["Post"] = post
	data["Content"] = template.HTML(post.Content)
	view.ExecuteTemplate(w, "index", data)
}
