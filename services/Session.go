package services

import (
	"os"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func GetStore() *sessions.CookieStore {
	if store == nil {
		store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	}
	store.Options = &sessions.Options{
		MaxAge:   60 * 20, //maksimum 20 dakika erişim ile kısıtladım
		HttpOnly: true,
		Path:     "/",
	}
	return store
}
