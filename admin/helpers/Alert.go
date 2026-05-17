package helpers

import (
	"fmt"
	"goblog/services"
	"net/http"
)

func SetAlert(w http.ResponseWriter, r *http.Request, message string) error {

	session, err := services.GetStore().Get(r, "go-alert")
	if err != nil {
		fmt.Println(err)
		return err
	}
	session.Options.MaxAge = 60 * 10
	session.AddFlash(message)

	return session.Save(r, w)

}

func GetAlert(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	session, err := services.GetStore().Get(r, "go-alert")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	data := make(map[string]interface{})
	flashes := session.Flashes()
	if len(flashes) > 0 {
		data["is_alert"] = true
		data["message"] = flashes[0]
	} else {
		data["is_alert"] = false
		data["message"] = nil
	}

	session.Save(r, w)

	return data
}
