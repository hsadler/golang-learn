package json_api

import(
	"fmt"
	"net/http"
	"encoding/json"
	"web_app/model"
)


// TODO: deprecate once finer grain response functions are written and
// dependencies are cut over
// func SendRes(w http.ResponseWriter, success bool, payload string) {
// 	apiResponse := &model.ApiResponse{
// 		Success: success,
// 		Payload: payload,
// 	}
// 	b, err := json.Marshal(apiResponse)
// 	if err != nil {
// 		fmt.Fprintf(w, err.Error())
// 		return
// 	}
// 	fmt.Fprintf(w, string(b))
// }

// TODO: implement more robust options for endpoint responses

func SendBoolApiResponse(
	w http.ResponseWriter,
	success bool,
	boolPayload bool,
) {
	// stub
}

func SendIntApiResponse(
	w http.ResponseWriter,
	success bool,
	intPayload int,
) {
	// stub
}

func SendStringApiResponse(
	w http.ResponseWriter,
	success bool,
	stringPayload string,
) {
	// stub
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



