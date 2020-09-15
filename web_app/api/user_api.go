package api

import (
	"net/http"
	"web_app/model"
	"web_app/service/jsonapi"
)

// GetUser : endpoint for getting mock user data
func GetUser(w http.ResponseWriter, req *http.Request) {
	user := &model.User{
		FirstName: "Harry",
		LastName:  "Sadler",
		Age:       1,
	}
	jsonapi.SendAPIResponse(w, true, user.GetAPIFormattedData())
}
