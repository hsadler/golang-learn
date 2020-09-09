package api

import (
	"fmt"
	"net/http"
	"web_app/service/jsonapi"
)

// Hello : sample "hello world" endpoint
func Hello(w http.ResponseWriter, req *http.Request) {
	success := true
	payload := "hello world"
	jsonapi.SendStringAPIResponse(w, success, payload)
}

// GetHeaders : endpoint to return request header information in a single string
func GetHeaders(w http.ResponseWriter, req *http.Request) {
	success := true
	payload := ""
	for name, headers := range req.Header {
		for _, h := range headers {
			payload += fmt.Sprintf("%s: %s\n", name, h)
		}
	}
	jsonapi.SendStringAPIResponse(w, success, payload)
}

// GetGETParams : endpoint to gather and return all GET params from request
func GetGETParams(w http.ResponseWriter, req *http.Request) {
	// stub
}

// GetPOSTParams : endpoint to gather and return all POST params from request
func GetPOSTParams(w http.ResponseWriter, req *http.Request) {
	// stub
}
