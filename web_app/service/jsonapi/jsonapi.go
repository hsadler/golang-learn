package jsonapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web_app/model"
)

// SendBoolAPIResponse : TODO
func SendBoolAPIResponse(
	w http.ResponseWriter,
	success bool,
	boolPayload bool,
) {
	apiResponse := &model.BoolAPIResponse{
		Success:     success,
		BoolPayload: boolPayload,
	}
	b, err := json.Marshal(apiResponse)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
}

// SendIntAPIResponse : TODO
func SendIntAPIResponse(
	w http.ResponseWriter,
	success bool,
	intPayload int,
) {
	apiResponse := &model.IntAPIResponse{
		Success:    success,
		IntPayload: intPayload,
	}
	b, err := json.Marshal(apiResponse)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
}

// SendStringAPIResponse : TODO
func SendStringAPIResponse(
	w http.ResponseWriter,
	success bool,
	stringPayload string,
) {
	apiResponse := &model.StringAPIResponse{
		Success:       success,
		StringPayload: stringPayload,
	}
	b, err := json.Marshal(apiResponse)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
}

// SendJSONAPIResponse : TODO
func SendJSONAPIResponse(
	w http.ResponseWriter,
	success bool,
	JSONPayload string,
) {
	apiResponse := &model.JSONAPIResponse{
		Success:     success,
		JSONPayload: JSONPayload,
	}
	b, err := json.Marshal(apiResponse)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
}
