package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"web_app/service/jsonapi"
)

// SamtryggSearch : TODO add docstring
func SamtryggSearch(w http.ResponseWriter, req *http.Request) {
	searchTerm := req.FormValue("search_term")
	getReqString := "https://www.samtrygg.se/RentalObject/SearchResult?search=" +
		url.QueryEscape(searchTerm)
	resp, err := http.Get(getReqString)
	if err != nil {
		jsonapi.SendStringAPIResponse(w, false, "problem executing search")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		jsonapi.SendStringAPIResponse(w, false, "problem executing search")
		return
	}
	var result interface{}
	json.Unmarshal(body, &result)
	jsonapi.SendAPIResponse(w, true, result)
}
