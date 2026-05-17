package middleware

import (
	"goblog/admin/helpers"
	"goblog/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func AdminMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if !services.ChechkUserIsAdmin(w, r) {
			helpers.SetAlert(w, r, "yetkisiz erişim!")
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}
		next(w, r, params)
	}
}

func ClientMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if !services.ChechkUserIsClient(w, r) {
			helpers.SetAlert(w, r, "Lütfen Giriş Yapınız!")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r, params)
	}
}
