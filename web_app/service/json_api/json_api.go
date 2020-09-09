package json_api

import(
	"fmt"
	"net/http"
	"encoding/json"
	"web_app/model"
)


func SendBoolApiResponse(
	w http.ResponseWriter,
	success bool,
	boolPayload bool,
) {
	apiResponse := &model.BoolApiResponse{
		Success: success,
		BoolPayload: boolPayload,
	}
	b, err := json.Marshal(apiResponse)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
}

func SendIntApiResponse(
	w http.ResponseWriter,
	success bool,
	intPayload int,
) {
	apiResponse := &model.IntApiResponse{
		Success: success,
		IntPayload: intPayload,
	}
	b, err := json.Marshal(apiResponse)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
}

func SendStringApiResponse(
	w http.ResponseWriter,
	success bool,
	stringPayload string,
) {
	apiResponse := &model.StringApiResponse{
		Success: success,
		StringPayload: stringPayload,
	}
	b, err := json.Marshal(apiResponse)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
}

func SendJsonApiResponse(
	w http.ResponseWriter,
	success bool,
	jsonPayload string,
) {
	apiResponse := &model.JsonApiResponse{
		Success: success,
		JsonPayload: jsonPayload,
	}
	b, err := json.Marshal(apiResponse)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
}

