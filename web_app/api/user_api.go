package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web_app/model"
	"web_app/service/jsonapi"
)

// GetUser :
func GetUser(w http.ResponseWriter, req *http.Request) {
	user := &model.User{Age: 1}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	success := true
	userJSON := string(b)
	jsonapi.SendJSONAPIResponse(
		w,
		success,
		userJSON,
	)
}
