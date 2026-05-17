package services

import (
	"net/http"
)

func SetUser(w http.ResponseWriter, r *http.Request, userID uint, mail string, role string) error {
	session, err := GetStore().Get(r, "system-user")
	if err != nil {
		return err
	}

	session.Values["userID"] = userID
	session.Values["mail"] = mail
	session.Values["role"] = role

	return session.Save(r, w)
}

func ChechkUserIsAdmin(w http.ResponseWriter, r *http.Request) bool {

	session, err := GetStore().Get(r, "system-user")
	if err != nil {
		return false
	}

	role := session.Values["role"]

	if role == "admin" {

		return true
	}

	return false

}

func ChechkUserIsClient(w http.ResponseWriter, r *http.Request) bool {
	session, err := GetStore().Get(r, "system-user")
	if err != nil {
		return false
	}

	role := session.Values["role"]
	//user := models.User{}.Get("email=?", mail)

	if role == "client" {
		return true
	}

	return false

}

func RemoveUser(w http.ResponseWriter, r *http.Request) error {

	session, err := GetStore().Get(r, "system-user")
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1

	return session.Save(r, w)

}
