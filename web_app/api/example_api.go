package api

import(
	"fmt"
	"net/http"
	"encoding/json"
	"web_app/model"
)


func Hello(w http.ResponseWriter, req *http.Request) {
	payload := "hello"
	apiResponse := &model.ApiResponse{Success: true, Payload: payload}
	b, err := json.Marshal(apiResponse)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(b))
}

// func headers(w http.ResponseWriter, req *http.Request) {
// 	for name, headers := range req.Header {
// 		for _, h := range headers {
// 			fmt.Fprintf(w, "%v: %v\n", name, h)
// 		}
// 	}
// }

