package api

import(
	"fmt"
	"net/http"
	"encoding/json"
	"web_app/model"
	"web_app/service/json_api"
)


func GetUser(w http.ResponseWriter, req *http.Request) {
	user := &model.User{Age:1}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	success := true
	userJson := string(b)
	json_api.SendJsonApiResponse(
		w,
		success,
		userJson,
	)
}
