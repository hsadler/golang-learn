package main

import (
	"net/http"
	"web_app/api"
)

func main() {

	var prepend string

	// example api
	prepend = "/api/example"
	http.HandleFunc(prepend+"/hello", api.Hello)
	http.HandleFunc(prepend+"/get_headers", api.GetHeaders)
	http.HandleFunc(prepend+"/get_get_params", api.GetGETParams)
	http.HandleFunc(prepend+"/get_post_params", api.GetPOSTParams)

	// user api
	prepend = "/api/user"
	http.HandleFunc(prepend+"/get_user", api.GetUser)

	// math api
	prepend = "/api/math"
	http.HandleFunc(prepend+"/add", api.Add)

	prepend = "/api/file_write"
	http.HandleFunc(prepend+"/write_value_to_file", api.WriteValueToFile)
	http.HandleFunc(prepend+"/read_value_from_file", api.ReadValueFromFile)

	// serve
	http.ListenAndServe(":8090", nil)

}
