package main

import (
	"net/http"
	"web_app/api"
)


func main() {

	http.HandleFunc("/hello", api.Hello)
	http.HandleFunc("/get_headers", api.GetHeaders)
	http.HandleFunc("/get_user", api.GetUser)

	http.ListenAndServe(":8090", nil)

}
