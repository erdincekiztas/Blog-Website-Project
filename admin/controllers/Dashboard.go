package controllers

import (
	"fmt"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"strconv"
	"time"

	"io"
	"os"

	"github.com/gosimple/slug"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll() //database sutunu yok

	data["Alert"] = helpers.GetAlert(w, r)

	err = view.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (dashboad Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println("\nyeni blog ekleme sayfası açılırken hata oluştu")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = view.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")

	//upload
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("picture-url")
	if err != nil {
		fmt.Println("fotoğraf alınırken sorun çıktı")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// upload sonu

	if title == "" || content == "" {

		time.Sleep(time.Millisecond * 1000)
		http.Redirect(w, r, "/admin/yeni-post-ekle", http.StatusSeeOther)
		return
	}

	models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		Content:     content,
		CategoryID:  categoryID,
		PictureUrl:  "/uploads/" + header.Filename,
	}.Add()

	helpers.SetAlert(w, r, "Blog Başarı İle Eklendi!")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	post := models.Post{}.Get(params.ByName("id"))

	//veriyi veritabanından silerken aynı zamanda uploads kısmından da silmek için
	if post.PictureUrl != "" {
		filepath := post.PictureUrl
		err := os.Remove(filepath[1:])
		fmt.Println(filepath)
		if err != nil {
			fmt.Println("Dosya Silinemedi:", err)
		}
	}
	post.Delete()
	helpers.SetAlert(w, r, "Blog Başarı İle Silindi!")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})

	data["Post"] = models.Post{}.Get(params.ByName("id"))

	view.ExecuteTemplate(w, "index", data)

}
func (dashboard Dashboard) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	post := models.Post{}.Get(params.ByName("id"))

	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")
	is_selected := r.FormValue("is_selected")
	var pictureUrl string

	if is_selected == "1" {

		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("picture-url")
		if err != nil {
			fmt.Println("fotoğraf alınırken sorun çıktı")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		pictureUrl = "/uploads/" + header.Filename
		fmt.Println(pictureUrl)

		// Sadece eski fotoğraf varsa sil
		if post.PictureUrl != "" {
			filepath := post.PictureUrl[1:]
			os.Remove(filepath)
			fmt.Println(filepath)
		}

	} else {
		pictureUrl = post.PictureUrl
	}

	post.Updates(models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		PictureUrl:  pictureUrl,
	})

	helpers.SetAlert(w, r, "Blog Başarı İle Güncellendi!")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
