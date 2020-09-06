package main

import (
	"net/http"
	"web_app/api"
)


func main() {

	var prepend string

	// example api
	prepend = "/api/example"
	http.HandleFunc(prepend + "/hello", api.Hello)
	http.HandleFunc(prepend + "/get_headers", api.GetHeaders)

	// user api
	prepend = "/api/user"
	http.HandleFunc(prepend + "/get_user", api.GetUser)

	// math api
	prepend = "/api/math"
	http.HandleFunc(prepend + "/add", api.Add)

	http.ListenAndServe(":8090", nil)

}
