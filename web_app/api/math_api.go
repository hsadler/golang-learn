package api

import (
	"net/http"
	"strconv"
	"web_app/service/jsonapi"
)

// Add : endpoint for adding two numbers and returning the result
func Add(w http.ResponseWriter, req *http.Request) {
	n1 := req.FormValue("n1")
	n2 := req.FormValue("n2")
	// alternate way of retrieving url query string params
	// from GET requests specifically
	// n1 := req.URL.Query().Get("n1")
	// n2 := req.URL.Query().Get("n2")
	n1int, _ := strconv.Atoi(n1)
	n2int, _ := strconv.Atoi(n2)
	addResult := n1int + n2int
	success := true
	jsonapi.SendIntAPIResponse(
		w,
		success,
		addResult,
	)
}
