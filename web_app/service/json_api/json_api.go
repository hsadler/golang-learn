package json_api

import(
	"fmt"
	"net/http"
	"encoding/json"
	"web_app/model"
)


func SendRes(w http.ResponseWriter, success bool, payload string) {
	apiResponse := &model.ApiResponse{
		Success: success,
		Payload: payload,
	}
	b, err := json.Marshal(apiResponse)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
}

// TODO: implement more robust options for enpoint responses
func SendJsonPayloadRes(w http.ResponseWriter, success bool, payload string) {}

