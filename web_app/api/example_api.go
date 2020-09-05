package api

import(
	"fmt"
	"net/http"
	"web_app/service/json_api"
)


func Hello(w http.ResponseWriter, req *http.Request) {
	success := true
	payload := "hello"
	json_api.SendRes(w, success, payload)
}

func GetHeaders(w http.ResponseWriter, req *http.Request) {
	success := true
	payload := ""
	for name, headers := range req.Header {
		for _, h := range headers {
			payload += fmt.Sprintf("%s: %s\n", name, h)
		}
	}
	json_api.SendRes(w, success, payload)
}

